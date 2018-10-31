package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"public.info/config"
	"public.info/server"
)

//log日志
func init() {
	config.ConfigLocalFilesystemLogger(config.Path, config.File, config.MaxAge, config.Interval)
}

//curl -X GET  "http://192.168.197.1:9999/ip" -d  '{"from":5,"size":65.23.2.1}'
func HandleIPServer(connect http.ResponseWriter, response *http.Request) {
	fmt.Println("start run  HandleIPServer")
	io.WriteString(connect, "welcome to ip page!\n")
	robots, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}

	server.DealQuerycmd(connect, string(robots))
	fmt.Printf("HandleIPServer running over:%v", string(robots))

}
func main() {
	//主动连接服务器
	http.HandleFunc("/ip", HandleIPServer)
	err := http.ListenAndServe(config.ADDRPORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
