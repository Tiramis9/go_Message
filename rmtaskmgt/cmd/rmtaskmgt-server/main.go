/****************************************
1. Create a server
2. Listening port
3. Registration service
4.Run the service
**************************************************************************/

package main

import (
	"github.com/sirupsen/logrus"

	Redio "huapu.info/rmcloud.exp/rmtaskmgt/server"
)

var (
	pointaddr = "127.0.0.1:8080"
	linsten   = "127.0.0.1:8888"
)

func main() {
	logger := logrus.WithField("func", "main")
	logger.Info("Welcome to rmtaskmgt-server")
	_, err := Redio.CreateRPCServer(pointaddr)
	if err != nil {
		logger.Fatalf("ponit Conenct err:%v\n", err)
		return
	}

	err = Redio.CreateListenConnect(linsten)
	if err != nil {
		logger.Fatalf("TCP Listen err:%v\n", err)
		return
	}

}
