package main

import (
	"flag"

	"github.com/golang/glog"
	cmd "public.info/cmd/task"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	cmd.Info()
}
