/****************************************************
CreateListenConnect The incoming parameter is a listening address

The registration service is executed inside the function

This function executes the run GRPC service

handleConnect his function is to get the requested address

RedioworkManage This function is for management status
*******************************************************/

package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC-Gateway"
	"huapu.info/rmcloud.exp/rmtaskmgt/pkg/util"
)

var (
	lineNumber = "line"
	skipNumber = 3
)

type client chan string

var (
	login   = make(chan client)
	message = make(chan string)
	logout  = make(chan client)
)

func RedioworkManage() {
	onlinClientMap := map[client]bool{}
	for {
		select {
		case client := <-login:
			onlinClientMap[client] = true
			//fmt.Println("login:%v\n", client)
		case client := <-logout:
			//fmt.Println("logout:%v\n", client)
			delete(onlinClientMap, client)

		case msg := <-message:
			fmt.Println("loginName:%v\n", msg)
			for cli, isOnline := range onlinClientMap {
				if isOnline == true {
					cli <- msg
				}
			}
		}
	}
}

func ReadFormMyClent(connect net.Conn, myCli client) {

	for msg := range myCli {
		connect.Write([]byte(msg + "\n"))
	}
}

func startServer() error {

	return nil
}

func handleConnect(connect net.Conn) {
	defer connect.Close()
	cliTimeOutTicker := time.NewTicker(time.Second * 10)
	revcData := make(chan interface{})
	removeClose := make(chan interface{})
	user := connect.RemoteAddr().String()
	cliect := make(client)

	go ReadFormMyClent(connect, cliect)

	login <- cliect
	message <- user + "Conect"

	go func() {
		input_scanf := bufio.NewScanner(connect)
		for input_scanf.Scan() {
			message <- user + ":" + input_scanf.Text()
			revcData <- struct{}{}
		}
		removeClose <- struct{}{}
	}()
	for {
		select {
		case <-revcData:
			cliTimeOutTicker.Stop()
			cliTimeOutTicker = time.NewTicker(time.Second * 10)
		case <-removeClose:
			logout <- cliect
			message <- user + "TXIT"
			return
		}
	}
}
func StartRedioServer(linsten string, pointaddr string) error {

	listen, err := net.Listen("tcp", linsten)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}
	RedioServer := newServer(pointaddr)

	if err = RedioServer.Serve(listen); err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}
	return err

}

func newServer(pointaddr string) *http.Server {
	grpcServer := RegisterWorking()
	gwmux, err := RegisterRadioFromEndpoint(pointaddr)
	if err != nil {
		panic(err)
	}

	httpmux := http.NewServeMux()
	httpmux.Handle("/", gwmux)
	httpmux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(httpmux)

	return &http.Server{
		Addr:    pointaddr,
		Handler: util.GrpcHandlerFunc(grpcServer, grpcServer),
	}
}

func CreateListenConnect(LinstenAddr string) error {
	rlog := logrus.WithField("server", "CreateListenConnect")
	rlog.Data[lineNumber] = FindCallerLine(skipNumber)
	rmiLinstenTasking, err := net.Listen("tcp", LinstenAddr)
	if err != nil {
		rlog.Printf("TCP Listen err:%v\n", err)
	}
	defer rmiLinstenTasking.Close()

	RadioProxy := MynewRadioProxy()
	Redioworking := grpc.NewServer()
	pb.RegisterRadioManageServer(Redioworking, RadioProxy)

	rlog.Printf("gRPC and https listen on: %s\n", rmiLinstenTasking)
	err = Redioworking.Serve(rmiLinstenTasking)
	if err != nil {
		rlog.Fatalln("Redioworking Failed..\n ")
	}
	/*
		for {
			connect, err := rmiLinstenTasking.Accept()
			if err != nil {
				rlog.Fatal(err)
				continue
			}
			go handleConnect(connect)

		} */
	return nil
}
