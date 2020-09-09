package grpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb_fetcher "github.com/wt-l00/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	pb_renderer "github.com/wt-l00/Hatena-Intern-2020/services/renderer-go/pb/renderer"
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

func Test_Server_Render(t *testing.T) {
	fetcherCli := &mockedFetcherClient{
		returnValue: "sample",
	}
	s := NewServer(fetcherCli)
	src := `[google](https://google.com/)`
	reply, err := s.Render(context.Background(), &pb_renderer.RenderRequest{Src: src})
	assert.NoError(t, err)
	assert.Equal(t, "<p><a href=\"https://google.com/\">google</a></p>\n", reply.Html)
}
