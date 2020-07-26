package main

import (
	"flag"
	userpb "grpc-samples/reverse-proxy-sample/api/userpb"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	REST_API_PORT = ":9997"
	GRPC_HOST     = "localhost:9998"
)

var userServiceEndpoint = flag.String("user_service_endpoint", GRPC_HOST, "Users sample")

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *userServiceEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(REST_API_PORT, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
