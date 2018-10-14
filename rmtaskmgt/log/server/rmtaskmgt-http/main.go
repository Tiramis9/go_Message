/*******************************************************************
1. Create a server
2. Listen on the port
3. Registration service
4.Run the service
*****************************************************************************************/
package main

import (
	"github.com/sirupsen/logrus"
	"huapu.info/rmcloud.exp/rmtaskmgt/server"
)

var (
	rlog       = logrus.WithField("pkg", "rmtaskmgt")
	listenaddr = "127.0.0.1:8999"
	pointaddr  = "127.0.0.1:8888"
)

func main() {
	logger := rlog.WithField("func", "main")
	logger.Info("Welcome to rmtaskmgt-server")
	//	go server.InitRadioProxy(pointaddr)
	err := server.StartRedioServer(listenaddr, pointaddr)
	if err != nil {
		logger.Printf("RedioTasking: %v\n", err)
		return
	}
}
