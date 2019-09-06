package auth

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
	google_rpc "istio.io/gogo-genproto/googleapis/google/rpc"
)

type server struct {
	fail      bool
	marshaler jsonpb.Marshaler
}

var _ v2.AuthorizationServer = &server{}

// New creates a new authorization server.
func New() v2.AuthorizationServer {
	_, fail := os.LookupEnv("FAIL")
	log.Printf("FAIL: %v", fail)
	return &server{fail: fail}
}

func (s *server) Check(ctx context.Context, req *v2.CheckRequest) (*v2.CheckResponse, error) {
	str, _ := s.marshaler.MarshalToString(req)
	log.Printf("%v", ctx)
	log.Println(str)

	time.Sleep(300 * time.Millisecond)
	code := int32(google_rpc.OK)
	if s.fail {
		code = int32(google_rpc.UNAUTHENTICATED)
	}

	return &v2.CheckResponse{
		HttpResponse: &v2.CheckResponse_OkResponse{
			OkResponse: &v2.OkHttpResponse{
				Headers: []*core.HeaderValueOption{
					{
						Append: &types.BoolValue{Value: false},
						Header: &core.HeaderValue{Key: "authorization", Value: "Bearer ok"},
					},
				},
			},
		},
		Status: &google_rpc.Status{
			Code: code,
		},
	}, nil
}
