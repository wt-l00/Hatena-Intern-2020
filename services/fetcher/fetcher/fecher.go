package fetcher

import (
	"context"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func scanHTML(node *html.Node) (string, bool) {
	if isTitleElement(node) {
		return node.FirstChild.Data, true
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if result, ok := scanHTML(c); ok {
			return result, ok
		}
	}
	return "", false
}

func getTitleFromHTML(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", false
	}
	return scanHTML(doc)
}

// Fetch title form url.
func Fetch(ctx context.Context, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return url, err
	}
	defer resp.Body.Close()

	if title, ok := getTitleFromHTML(resp.Body); ok {
		return title, nil
	}
	// titleが取得できない場合はurlを返しておきます
	return url, nil
}
