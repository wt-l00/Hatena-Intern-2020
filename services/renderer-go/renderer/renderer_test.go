package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Render_Htag(t *testing.T) {
	htagTests := []struct {
		actual   string
		expected string
	}{
		{"# h1", "<h1>h1</h1>\n"},
		{"## h2", "<h2>h2</h2>\n"},
		{"### h3", "<h3>h3</h3>\n"},
		{"#### h4", "<h4>h4</h4>\n"},
		{"##### h5", "<h5>h5</h5>\n"},
		{"###### h6", "<h6>h6</h6>\n"},
	}

	for _, htag := range htagTests {
		html, err := Render(context.Background(), htag.actual)
		assert.NoError(t, err)
		assert.Equal(t, htag.expected, html)
	}
}

func Test_Render_Link(t *testing.T) {
	src := `[Google](https://google.com)`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, "<p><a href=\"https://google.com\">Google</a></p>\n", html)
}

func Test_Render_List(t *testing.T) {
	src := `
- list1
- list2
- list3`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<ul>
<li>list1</li>
<li>list2</li>
<li>list3</li>
</ul>
`, html)
}

func Test_Render_Commentout(t *testing.T) {
	src := `//TODO: something//`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, "<!-- TODO: something -->", html)
}
