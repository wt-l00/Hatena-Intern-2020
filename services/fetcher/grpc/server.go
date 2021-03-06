package grpc

import (
	"context"

	"github.com/wt-l00/Hatena-Intern-2020/services/fetcher/fetcher"
	pb "github.com/wt-l00/Hatena-Intern-2020/services/fetcher/pb/fetcher"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.FetcherServer に対する実装
type Server struct {
	pb.UnimplementedFetcherServer
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer() *Server {
	return &Server{}
}

// Fetch は受け取った url から title を取得
func (s *Server) Fetch(ctx context.Context, in *pb.FetcherRequest) (*pb.FetcherReply, error) {
	title, err := fetcher.Fetch(ctx, in.Src)
	if err != nil {
		return nil, err
	}
	return &pb.FetcherReply{Title: title}, nil
}
