package server

import (
	"net"

	"github.com/sirupsen/logrus"
)

func QueryWorking(lis *net.Listener) error {

	rlog := logrus.WithField("server", "QueryWorking")
	rlog.Data[lineNumber] = FindCallerLine(skipNumber)

	rlog.Info("hehh")
	return nil
}

// 建立rpc连接
func temp() {
	/*
		conn, err := grpc.Dial(*serverAddr)
		if err != nil {

		}
		defer conn.Close()
		client := pb.NewRouteGuideClient(conn)
		feature, err := client.GetFeature(context.Background(), &pb.Point{409146138, -746188906})
		if err != nil {

		}
	*/
}
