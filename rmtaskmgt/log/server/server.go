package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC" // 引入编译生成的包
	comm "huapu.info/rmcloud.exp/rmtaskmgt/comm"
	// 引入编译生成的包
)

// Create a new instance of the logger. You can have any number of instances.
var (
	Log_Path = string("../log/")
	Log_File = string("std.log")
	MaxAge   = time.Hour * 24
	maxSlice = time.Second * 2000
)

func init() {
	// ConfigLocalFilesystemLogger 第一个参数PATH，第二个参数File,第三个参数文件为最大保存时间，第四个参数为切割日志时间间隔
	comm.ConfigLocalFilesystemLogger(Log_Path, Log_File, MaxAge, maxSlice)
}

var (
	Address      = "192.168.140.30:8080"
	echoEndpoint = flag.String("echo_endpoint", Address, "endpoint of YourService")
	CertPemPath  = "src/huapu.info/rmcloud.exp/rmtaskmgt/keys"
	CertKeyPath  = "src/huapu.info/rmcloud.exp/rmtaskmgt/keys"
)

func LinstenServe() (err error) {
	conn, err := net.Listen("tcp", Address)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}
	/*tlsConfig := util.GetTLSConfig(CertPemPath, CertKeyPath)
	srv := createInternalServer(conn, tlsConfig)
	/*
	//log.Printf("gRPC and https listen on: %s\n", Address)
	comm.Log.Printf("createInternalServer")
	srv.Serve(conn)

		if err = srv.Serve(tls.NewListener(conn, tlsConfig)); err != nil {
			log.Printf("ListenAndServe: %v\n", err)
		}*/
	comm.Log.Printf("srv.Serve")
	return err
}
func createInternalServer(conn net.Listener, tlsConfig *tls.Config) *http.Server {
	comm.Log.Printf("G etTLSConfig")
	// var opts []grpc.ServerOption

	// grpc server
	/* creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		log.Printf("Failed to create server TLS credentials %v", err)
	}

	//opts = append(opts, grpc.Creds(creds))
	//server := grpc.NewServer(opts...)
	*/
	// 注册grpc server
	server := grpc.NewServer()
	gw.RegisterManageServeServer(server, NewManageServe())
	fmt.Println("linsten" + Address)
	// 注册 关联组件

	/* dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, "123")
	if err != nil {
		log.Printf("Failed to create client TLS credentials %v", err)
	}
	// dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	*/
	//	comm.Log.Printf("G etTLSConfig")
	ctx := context.Background()
	dop := []grpc.DialOption{grpc.WithInsecure()}
	gwmux := runtime.NewServeMux()

	comm.Log.Printf("DialOption")
	// 注册 grpc-gateway pb
	if err := gw.RegisterManageServeHandlerFromEndpoint(ctx, gwmux, *echoEndpoint, dop); err != nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}
	comm.Log.Printf("RegisterManageServeHandlerFromEndpoint")
	// http服务
	mux := http.NewServeMux()
	comm.Log.Printf("NewServeMux")
	mux.Handle("/", gwmux)
	//	return http.ListenAndServe("192.168.140.30:8080", mux)
	return &http.Server{
		Addr:    *echoEndpoint,
		Handler: GrpcHandlerFunc(server, mux),
		// TLSConfig: tlsConfig,
	}
}
func main() {
	LinstenServe()
}
func GrpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	if otherHandler == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			grpcServer.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
