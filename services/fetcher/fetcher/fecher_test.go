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
		{"https://github.com", "The world’s leading software development platform · GitHub"},
	}
	for _, url := range fetchTests {
		title, err := Fetch(context.Background(), url.url)
		assert.NoError(t, err)
		assert.Equal(t, title, url.expectedTitle)
	}
}
