package renderer

import (
	"bytes"
	"context"

	commentout "github.com/wt-l00/goldmark-commentout"
	"github.com/yuin/goldmark"
)

// Render は受け取った文書をHTMLとして返す
func Render(ctx context.Context, src string) (string, error) {
	html, err := ConvertMd(src)
	if err != nil {
		return "", err
	}
	return html, nil
}

// ConvertMd は受け取った文書（markdown）を HTMLに変換する
func ConvertMd(src string) (string, error) {
	var buf bytes.Buffer
	var md = goldmark.New(
		goldmark.WithExtensions(
			commentout.Commentout,
		),
	)
	if err := md.Convert([]byte(src), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
