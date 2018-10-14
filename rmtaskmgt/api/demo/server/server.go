/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	rmon_service "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC/rpc"
	"huapu.info/rmcloud.exp/rmtaskmgt/api/demo/protocol"
)

const (
	RecPort  = ":61000"
	SendPort = ":50053"

	OK          = 0
	ERROR_VALUE = -1
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

// 混合类型的struct
type ComplexData struct {
	ID          int64
	TimeStamp   int64
	PackageNum  int64
	PackageSize int64
	//	PackageData   []byte
}

func send(conn net.Conn, ID int64, PackageSize int64, PackageSpeed int64) int {
	/*
		for i := 0; i < 100; i++ {
			session := GetSession()
			words := "{\"ID\":" + strconv.Itoa(i) + "\",\"Session\":" + session + "2015073109532345\",\"Meta\":\"golang\",\"Content\":\"message\"}"
			conn.Write(protocol.Enpack([]byte(words)))
		}
	*/

	//get time stamp , Time type
	timestamp := time.Now().UnixNano()

	fmt.Println(timestamp)

	PackageNum := PackageSpeed / PackageSize

	if 0 == PackageNum {
		PackageNum = 1
	}

	// make package size
	DataBytes := make([]byte, PackageSize)

	var i int64
	for i = 0; i < PackageNum; i++ {
		// ID  +  时间戳  +  包个数  +   包大小 + 包的数据
		words := strconv.FormatInt(ID, 10) + strconv.FormatInt(timestamp, 10) + strconv.FormatInt(PackageSize, 10) + strconv.FormatInt(PackageNum, 10)
		//words := strconv.FormatInt(ID, 10) + ":" + strconv.FormatInt(timestamp, 10) + ":" +strconv.FormatInt(PackageSize, 10) + ":" + strconv.FormatInt(PackageNum, 10) + ":"
		_, err := conn.Write(protocol.Enpack([]byte(words)))
		_, err1 := conn.Write(protocol.Enpack(DataBytes))

		fmt.Println(words)
		fmt.Println(DataBytes)

		if err != nil || err1 != nil {
			return ERROR_VALUE
		}

	}

	fmt.Println("send over")

	defer conn.Close()

	return OK
}

func GetSession() string {
	gs1 := time.Now().UnixNano()
	gs2 := strconv.FormatInt(gs1, 10)
	return gs2
}

func (s *server) SendMessage(ctx context.Context, in *rmon_service.CustomerRequest) (*rmon_service.ServiceReply, error) {
	fmt.Println(in.ID)
	fmt.Println(in.PackageSize)
	fmt.Println(in.PackageSpeed)

	if 0 == in.ID || 0 == in.PackageSize || 0 == in.PackageSpeed {
		return &rmon_service.ServiceReply{ResultInfo: rmon_service.ReturnValue_VALUE_ERRPR}, nil
	}

	/*
		//get time stamp , Time type
		timestamp := time.Now().UnixNano()

		fmt.Println(timestamp)

		PackageNum := in.PackageSpeed / in.PackageSize

		if 0 == PackageNum {
			PackageNum = 1
		}
	*/

	/*
		DataBytes := make([]byte, in.PackageSize)
		fmt.Println(len(DataBytes))

		fmt.Println("data:")
		fmt.Println(in.ID)
		fmt.Println(timestamp)
		fmt.Println(PackageNum)
		fmt.Println(in.PackageSize)

		HEADBytes := make([]byte, 32)

		var temptime uint64 = uint64(timestamp)

		binary.BigEndian.PutUint32(HEADBytes[0:], in.ID)
		binary.BigEndian.PutUint64(HEADBytes[8:], temptime)
		binary.BigEndian.PutUint64(HEADBytes[16:], PackageNum)
		binary.BigEndian.PutUint64(HEADBytes[24:], in.PackageSize)

		var buffer bytes.Buffer
		buffer.Write(HEADBytes)
		buffer.Write(DataBytes)

		msg := buffer.Bytes()
		fmt.Println(string(msg))
	*/
	//go net client
	server := "localhost:61000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return &rmon_service.ServiceReply{ResultInfo: rmon_service.ReturnValue_VALUE_ERRPR}, nil
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return &rmon_service.ServiceReply{ResultInfo: rmon_service.ReturnValue_VALUE_ERRPR}, nil
	}
	fmt.Println("connect success")

	temp := send(conn, in.ID, in.PackageSize, in.PackageSpeed)
	if temp < 0 {
		return &rmon_service.ServiceReply{ResultInfo: rmon_service.ReturnValue_VALUE_ERRPR}, nil
	}

	/*
		address := "localhost"
		conn, err := net.Dial("tcp", address+SendPort)
		if err != nil {
			return &rmon_service.ServiceReply{ResultInfo: rmon_service.ReturnValue_VALUE_ERRPR}, nil
		}

		var i uint64
		for i = 0; i < PackageNum; i++ {
			_, err := conn.Write(msg)

			fmt.Println(msg)

			if err != nil {
				return &rmon_service.ServiceReply{ResultInfo: rmon_service.ReturnValue_VALUE_ERRPR}, nil
			}

		}

		if err != nil {
			return &rmon_service.ServiceReply{ResultInfo: rmon_service.ReturnValue_VALUE_ERRPR}, nil
		}

		defer conn.Close()
	*/
	return &rmon_service.ServiceReply{ResultInfo: rmon_service.ReturnValue_SUCCESS}, nil
}

func main() {
	lis, err := net.Listen("tcp", RecPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rmon_service.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("6")
}
