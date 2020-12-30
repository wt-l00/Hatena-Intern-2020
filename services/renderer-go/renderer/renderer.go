package renderer

import (
	"bytes"
	"context"
	"sync"

	pb_fetcher "github.com/wt-l00/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	commentout "github.com/wt-l00/goldmark-commentout"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type autoTitleLinker struct {
	ctx        context.Context
	fetcherCli pb_fetcher.FetcherClient
}

// Render は受け取った文書をHTMLとして返す
func Render(ctx context.Context, src string, fetcherClient pb_fetcher.FetcherClient) (string, error) {
	html, err := ConvertMdtoHTML(ctx, src, fetcherClient)
	if err != nil {
		return "", err
	}
	return html, nil
}

// ConvertMdtoHTML は受け取った文書（markdown）を HTMLに変換する
func ConvertMdtoHTML(ctx context.Context, src string, fetcherCli pb_fetcher.FetcherClient) (string, error) {
	var buf bytes.Buffer
	var md = goldmark.New(
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&autoTitleLinker{fetcherCli: fetcherCli, ctx: ctx}, 999),
			),
		),
		goldmark.WithExtensions(
			commentout.Commentout,
		),
	)
	if err := md.Convert([]byte(src), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (l *autoTitleLinker) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	// url と nodeの関係を取り持つ．1 url : n node
	urlNodes := make(map[string][]*ast.Link)
	// url と titleの関係を取り持つ． 1 url : 1 title
	var urlTitle sync.Map

	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if node, ok := node.(*ast.Link); ok && entering && node.ChildCount() == 0 {
			urlNodes[string(node.Destination)] = append(urlNodes[string(node.Destination)], node)
		}
		return ast.WalkContinue, nil
	})

	wg := sync.WaitGroup{}

	for url := range urlNodes {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			title := fetchTitle(l.ctx, l.fetcherCli, url)
			urlTitle.Store(url, title)
		}(url)
	}
	wg.Wait()

	for url, nodes := range urlNodes {
		for _, node := range nodes {
			if title, ok := urlTitle.Load(url); ok {
				node.AppendChild(node, ast.NewString([]byte(title.(string))))
			}
		}
	}
}

// fetchTitle は FetcherClient を使用
func fetchTitle(ctx context.Context, fetcherCli pb_fetcher.FetcherClient, url string) string {
	reply, err := fetcherCli.Fetch(ctx, &pb_fetcher.FetcherRequest{Src: url})
	if err != nil {
		return url
	}
	return reply.Title
}
