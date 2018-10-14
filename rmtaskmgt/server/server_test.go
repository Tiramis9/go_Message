package server

import (
	"github.com/sirupsen/logrus"
)

var (
	addr = "127.0.0.1:"
	port = "8080"
)

func startLinstenTasking_Test() error {

	rlog := logrus.WithField("server", "StartTaskLinsten")
	rlog.Data[lineNumber] = FindCallerLine(skipNumber)
	//	if RedioListen, err := RedioListen("nil"); err != nil {
	rlog.Fatalf("StartLinstenTasking failed%v", "err")
	//	}
	return nil
}
func main() {
	rlog := logrus.WithField("server", "StartTaskLinsten")
	rlog.Data[lineNumber] = FindCallerLine(skipNumber)

	if err := startLinstenTasking_Test(); err != nil {
		rlog.Fatalf("StartLinstenTasking failed:%s", err)
	}
	if err := QueryWorking(nil); err != nil {
		rlog.Fatalf("QueryWorking failed:%s", err)
	}
}
