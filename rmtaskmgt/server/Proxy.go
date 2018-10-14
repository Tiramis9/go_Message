/***********************************************************
This is the service that implements Tasking
Assign tasks to the pointed endpoint
Assign the caller a task ID
************************************************************************/
package server

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	workProxy "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC-Gateway"
)

type RadioProxy struct {
	WorkID   string
	ponintid string
}

func (radioProxy *RadioProxy) Tasking(ctx context.Context, in *workProxy.WorkRequest) (*workProxy.WorkReply, error) {
	Result := new(workProxy.WorkReply)

	if in.StartFrequency == 0 || in.PackageSize == 0 {
		return &workProxy.WorkReply{WorkID: "The parameter cannot be empty"}, nil
	}
	connection, err := grpc.Dial(radioProxy.WorkID, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect error: %s", err.Error())
		return &workProxy.WorkReply{WorkID: "connection nil"}, err
	}
	defer connection.Close()
	//pointSer := point.NewGreeterClient(connection)
	/*
		pointSer, err := point.SendMessage(ctx, &point.CustomerRequest{ID: in.UserID, PackageSize: in.PackageSize, StartFrequency: in.StartFrequency})
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not greet: %s", err.Error())
			return &workProxy.WorkReply{WorkID: "connection nil"}, err
		}*/
	Result.WorkID = UniqueId()
	return Result, nil
}

func MynewRadioProxy() *RadioProxy {
	return new(RadioProxy)
}
func InitRadioProxy(pointaddr string) error {
	rlog := logrus.WithField("server", "InitRadioProxy")
	if pointaddr == " " {
		rlog.Info("param  pointaddr can't not ")
		return nil
	}
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		rlog.Fatal(err)
	}
	defer conn.Close()

	client := workProxy.NewRadioManageClient(conn)
	reply, err := client.Tasking(context.Background(), &workProxy.WorkRequest{MessageId: "helllo"})
	if err != nil {
		rlog.Fatal(err)

	}
	fmt.Println(reply.WorkID)

	return nil
}

func CreateRPCServer(pointaddr string) (*RadioProxy, error) {
	rlog := logrus.WithField("server", "CreateRPCServer")
	rlog.Info("Create RPCServer SUCCEED ")
	if pointaddr == " " {
		rlog.Fatal("CreateRPCServer failed..\n")
		return nil, nil
	}

	gRPCpoint := &RadioProxy{WorkID: UniqueId(), ponintid: pointaddr}
	return gRPCpoint, nil
}

func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
