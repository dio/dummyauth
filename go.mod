module github.com/dio/dummyauth

go 1.12

require (
	github.com/envoyproxy/go-control-plane v0.9.2

	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.3.2
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	golang.org/x/sys v0.0.0-20190830142957-1e83adbbebd0 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200113173426-e1de0a7b01eb
	google.golang.org/grpc v1.26.0
	istio.io/gogo-genproto v0.0.0-20190826122855-47f00599b597
)

replace github.com/envoyproxy/go-control-plane => github.com/dio/456bf9a-go-control-plane v0.0.0-20200114124407-9d640d104b7e
