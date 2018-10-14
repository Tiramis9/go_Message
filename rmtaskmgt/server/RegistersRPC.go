/***********************************************************

First register the GRPC service  BindManagesTask returns grpc.ServerAddrs
The endpoint service is then registered RegisterRadioFromEndpoint returns
http.Handler Create with this handle gateway Add http.Handler processing method
The last returned is the listening address and the GRPC handle

context.Background：返回一个非空的空上下文。它没有被注销，没有值，没有过期时间。
它通常由主函数、初始化和测试使用，并作为传入请求的顶级上下文
************************************************************************/
package server

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC-Gateway"
)

func RegisterRedioServer(bindaddr string) (*http.Server, error) {
	rlog := logrus.WithField("server", "RegisterRedioServer")
	if bindaddr == " " {
		rlog.Fatalf("newRegisterHTTPHandler")
		return nil, nil
	}
	RegisterTask := RegisterWorking()

	HandlerMux, err := RegisterRadioFromEndpoint(bindaddr)
	if err != nil {
		rlog.Fatalf("newRegisterHTTPHandler")
		return nil, err
	}
	hTTPServeMuxEntry, err := AddHTTPHandler(HandlerMux)
	if err != nil {
		rlog.Fatalf("CreateHTTPHandler Fa")
		return nil, err
	}
	return &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: GrpcHandlerFunc(RegisterTask, hTTPServeMuxEntry),
	}, nil
}

func AddHTTPHandler(HandlerMux http.Handler) (*http.ServeMux, error) {
	hTTPServeMuxEntry := http.NewServeMux()
	hTTPServeMuxEntry.Handle("/", HandlerMux)
	hTTPServeMuxEntry.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(hTTPServeMuxEntry)
	return hTTPServeMuxEntry, nil
}
func RegisterRadioFromEndpoint(bindaddr string) (http.Handler, error) {
	rlog := logrus.WithField("server", "RegisterRadioFromEndpoint")
	if bindaddr == " " {
		rlog.Fatal("RegisterRadioFromEndpoint bindaddr not nil", bindaddr)
		return nil, nil
	}
	endpointService := flag.String("endpoint", bindaddr, "endpoint of YourService")
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	GatewayMuxer := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	result := pb.RegisterRadioManageHandlerFromEndpoint(context, GatewayMuxer, *endpointService, opts)
	if result != nil {
		rlog.Fatal("RegisterRadioManageHandlerFromEndpoint", result)
		return nil, result
	}
	rlog.Infof("RegisterRadioManageHandlerFromEndpoint result %v\n\n", result)
	return GatewayMuxer, nil
}
func RegisterWorking() *grpc.Server {
	rlog := logrus.WithField("server", "BindManagesTask")
	rlog.Infof("RegisterRadioManageServer:\n")
	RadioProxy := MynewRadioProxy()
	Redioworking := grpc.NewServer()
	pb.RegisterRadioManageServer(Redioworking, RadioProxy)
	return Redioworking
}
