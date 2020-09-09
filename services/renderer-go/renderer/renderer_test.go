package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb_fetcher "github.com/wt-l00/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	"google.golang.org/grpc"
)

type fetcherClientInterface interface {
	Fetch(ctx context.Context, in *pb_fetcher.FetcherRequest, opts ...grpc.CallOption) (*pb_fetcher.FetcherReply, error)
}

type mockedFetcherClient struct {
	returnValue string
}

func (m *mockedFetcherClient) Fetch(ctx context.Context, in *pb_fetcher.FetcherRequest, opts ...grpc.CallOption) (*pb_fetcher.FetcherReply, error) {
	return &pb_fetcher.FetcherReply{Title: m.returnValue}, nil
}

func Test_Render_Htag(t *testing.T) {
	fetcherCli := &mockedFetcherClient{
		returnValue: "sample",
	}

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
		html, err := Render(context.Background(), htag.actual, fetcherCli)
		assert.NoError(t, err)
		assert.Equal(t, htag.expected, html)
	}
}

func Test_Render_Link(t *testing.T) {
	fetcherCli := &mockedFetcherClient{
		returnValue: "sample",
	}
	src := `[Google](https://google.com)`
	html, err := Render(context.Background(), src, fetcherCli)
	assert.NoError(t, err)
	assert.Equal(t, "<p><a href=\"https://google.com\">Google</a></p>\n", html)
}

func Test_Render_List(t *testing.T) {
	fetcherCli := &mockedFetcherClient{
		returnValue: "sample",
	}
	src := `
- list1
- list2
- list3`
	html, err := Render(context.Background(), src, fetcherCli)
	assert.NoError(t, err)
	assert.Equal(t, `<ul>
<li>list1</li>
<li>list2</li>
<li>list3</li>
</ul>
`, html)
}

func Test_Render_Commentout(t *testing.T) {
	fetcherCli := &mockedFetcherClient{
		returnValue: "Example Domain",
	}
	src := `//TODO: something//`
	html, err := Render(context.Background(), src, fetcherCli)
	assert.NoError(t, err)
	assert.Equal(t, "<!-- TODO: something -->", html)
}

func Test_Render_Autolink(t *testing.T) {
	fetcherCli := &mockedFetcherClient{
		returnValue: "Example Domain",
	}
	src := `[](https://example.com)`
	html, err := Render(context.Background(), src, fetcherCli)
	assert.NoError(t, err)
	assert.Equal(t, "<p><a href=\"https://example.com\">Example Domain</a></p>\n", html)
}
