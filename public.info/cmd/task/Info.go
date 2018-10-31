package task

import (
	"net"

	"public.info/config"
	"public.info/server"
)

//log日志
func init() {
	config.ConfigLocalFilesystemLogger(config.Path, config.File, config.MaxAge, config.Interval)
}

//开启监听服务
func Info1() {
	log := config.Log.WithField("package", "Task")
	log.Info("start running Info")
	go server.ManagerInfo()
	listen, err := net.Listen("tcp", config.ADDRPORT1)
	if err != nil {
		log.Println("net.Listen err = ", err)
		return
	}
	defer listen.Close()
	for {
		connect, err := listen.Accept()
		if err != nil {
			log.Println("listener.Accept err = ", err)
			continue
		}
		go server.HandleConnect(connect)
	}
}
func Info() {
	log := config.Log.WithField("package", "Task")
	log.Info("start running Info")
	go server.ManagerInfo()
	listen, err := net.Listen("tcp", config.ADDRPORT1)
	if err != nil {
		log.Println("net.Listen err = ", err)
		return
	}
	defer listen.Close()
	for {
		connect, err := listen.Accept()
		if err != nil {
			log.Println("listener.Accept err = ", err)
			continue
		}
		go server.HandleConnect(connect)
	}
}
