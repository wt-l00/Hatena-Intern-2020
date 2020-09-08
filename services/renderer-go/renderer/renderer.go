package renderer

import (
	"bytes"
	"context"

	"github.com/yuin/goldmark"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(src), &buf); err != nil {
		panic(err)
	}
	return buf.String(), nil
}
