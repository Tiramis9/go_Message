package main

import (
	"code/grpc/Node_Master/comm"
	"time"

	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var (
	Log_Path = string("D:/Vs_code/go_Setlup/bin/src/code/grpc/Node_Master/log/")
	Log_File = string("std.log")
	MaxAge   = time.Hour * 24
	maxSlice = time.Second * 2000
)

func init() {
	// ConfigLocalFilesystemLogger 第一个参数PATH，第二个参数File,第三个参数文件为最大保存时间，第四个参数为切割日志时间间隔
	comm.ConfigLocalFilesystemLogger(Log_Path, Log_File, MaxAge, maxSlice)
}
func main() { // example

	comm.Log.WithFields(logrus.Fields{
		"test": "test",
		"size": 10,
	}).Info("A group of walrus emerges from the ocean")
	comm.Log.Warning("warnign.")
	//	comm.Log.Fatal("Fatal")
	//设置输出格式
	// comm.Log.SetFormatter(&logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})
	comm.Log.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: false, TimestampFormat: "2006-01-02 15:04:05.000"})
	comm.Log.Info("Info")
	comm.Log.Debug("Debug")
	comm.Log.Println("test Log")
}
