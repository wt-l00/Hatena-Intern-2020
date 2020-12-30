package fetcher

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fetch_Title(t *testing.T) {
	fetchTests := []struct {
		url           string
		expectedTitle string
	}{
		{"https://www.google.com/search/about/", "Google Search - Stay in the Know with Your Google App"},
		{"https://books.google.com/books/about/ ", "Google Books"},
		{"https://hatenablog.com/api/", "https://hatenablog.com/api/"},
        {"https://example.com", "Example Domain"},
	}

	for _, fetchTest := range fetchTests {
		title, err := Fetch(context.Background(), fetchTest.url)
		assert.NoError(t, err)
		assert.Equal(t, fetchTest.expectedTitle, title)
	}
}
