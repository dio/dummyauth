package main

import (
	"log"
	"net"

	v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"google.golang.org/grpc"

	"github.com/dio/dummyauth/pkg/auth"
)

func main() {
	grpcServer := grpc.NewServer()
	v2.RegisterAuthorizationServer(grpcServer, auth.New())

	l, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Listening on tcp://localhost:3001")
	grpcServer.Serve(l)
}
