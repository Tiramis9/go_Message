package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC-Gateway"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8999", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRadioManageClient(conn)
	reply, err := client.Tasking(context.Background(), &pb.WorkRequest{MessageId: "helllo"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.WorkID)
}
