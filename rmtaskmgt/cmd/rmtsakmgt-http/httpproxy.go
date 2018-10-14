package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pw "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC-Gateway"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:8888", "endpoint of YourService")
)

func redaiorun() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pw.RegisterRadioManageHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := redaiorun(); err != nil {
		glog.Fatal(err)
	}
}
