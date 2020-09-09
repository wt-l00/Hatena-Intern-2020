package grpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/wt-l00/Hatena-Intern-2020/services/fetcher/pb/fetcher"
)

func Test_Server_Fetch_Title(t *testing.T) {
	s := NewServer()
	url := "https://example.com/"
	extected := "Example Domain"
	reply, err := s.Fetch(context.Background(), &pb.FetcherRequest{Src: url})
	assert.NoError(t, err)
	assert.Equal(t, extected, reply.Title)
}
