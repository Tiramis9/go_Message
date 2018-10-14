package main

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	pb "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC" // 引入编译生成的包
	rpc "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC/rpc"
)

var (
	address = "localhost:61000"
)

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/// 定义HTTPService并实现约定的接口
type manageServe struct {
}

// n\NewManageServe  HTTP服务
func NewManageServe() pb.ManageServeServer {
	return new(manageServe)
}

// SayHeManageMessagello 实现manage服务接口
func (m *manageServe) ManageMessage(ctx context.Context, in *pb.RmsmomRequest) (*pb.RmsmomReply, error) {
	resp := new(pb.RmsmomReply)
	// Set up a connection to the server.
	resp.Info = pb.Status_SUCCESS
	resp.TASKID = UniqueId()

	log.Println(in.UserId)
	log.Println(in.PackageSize)
	log.Println(in.PackageSpeed)
	fmt.Fprintf(os.Stderr, "Connect SUCCESS%s", "connect")
	if in.PackageSpeed == 0 || in.PackageSize == 0 {
		return &pb.RmsmomReply{Info: pb.Status_VALUE_ERRPR}, nil
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect error: %s", err.Error())
		return &pb.RmsmomReply{Info: pb.Status_VALUE_ERRPR}, err
	}
	defer conn.Close()

	c := rpc.NewGreeterClient(conn)
	r, err := c.SendMessage(ctx, &rpc.CustomerRequest{ID: in.UserId, PackageSize: in.PackageSize, PackageSpeed: in.PackageSpeed})
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not greet: %s", err.Error())
		return &pb.RmsmomReply{Info: pb.Status_VALUE_ERRPR}, err
	}
	log.Printf("Result:,%s", r.ResultInfo)
	resp.Info = pb.Status_SUCCESS
	resp.TASKID = UniqueId()
	return resp, nil
}
