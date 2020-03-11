package auth

import (
	"context"
	"log"
	"os"
	"strings"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
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

	statusCode := int32(code.Code_OK)
	if s.fail || strings.Contains(str, `["insert"]`) {
		statusCode = int32(code.Code_PERMISSION_DENIED)
	}
	return &v2.CheckResponse{
		HttpResponse: &v2.CheckResponse_OkResponse{
			OkResponse: &v2.OkHttpResponse{
				Headers: []*envoy_api_v2_core.HeaderValueOption{
					{
						Append: &wrappers.BoolValue{Value: false},
						Header: &envoy_api_v2_core.HeaderValue{Key: "authorization", Value: "Bearer ok"},
					},
				},
			},
		},
		Status: &status.Status{
			Code: statusCode,
		},
	}, nil
}
