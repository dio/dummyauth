package auth

import (
	"context"
	"log"
	"time"

	v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/gogo/googleapis/google/rpc"
	"github.com/golang/protobuf/jsonpb"
)

type server struct {
	marshaler jsonpb.Marshaler
}

var _ v2.AuthorizationServer = &server{}

// New creates a new authorization server.
func New() v2.AuthorizationServer {
	return &server{}
}

func (s *server) Check(ctx context.Context, req *v2.CheckRequest) (*v2.CheckResponse, error) {
	str, _ := s.marshaler.MarshalToString(req)
	log.Printf("%v", ctx)
	log.Println(str)

	time.Sleep(300 * time.Millisecond)

	return &v2.CheckResponse{
		Status: &rpc.Status{
			Code:    int32(rpc.OK),
			Message: "",
		},
	}, nil
}
