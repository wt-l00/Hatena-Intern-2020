package grpc

import (
	"context"

	pb_fetcher "github.com/wt-l00/Hatena-Intern-2020/services/renderer-go/pb/fetcher"
	pb_renderer "github.com/wt-l00/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	"github.com/wt-l00/Hatena-Intern-2020/services/renderer-go/renderer"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.RendererServer に対する実装
type Server struct {
	pb_renderer.UnimplementedRendererServer
	pb_fetcher.FetcherClient
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer(fetcherClient pb_fetcher.FetcherClient) *Server {
	return &Server{
		FetcherClient: fetcherClient,
	}
}

// Render は受け取った文書を HTML に変換する
func (s *Server) Render(ctx context.Context, in *pb_renderer.RenderRequest) (*pb_renderer.RenderReply, error) {
	html, err := renderer.Render(ctx, in.Src, s.FetcherClient)
	if err != nil {
		return nil, err
	}
	return &pb_renderer.RenderReply{Html: html}, nil
}
