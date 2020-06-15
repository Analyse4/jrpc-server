package main

import (
	"fmt"
	"os"

	"github.com/Analyse4/jrpc-server/jRPC"
	jruntime "github.com/Analyse4/jrpc-server/jRPC/runtime"
	"github.com/Analyse4/jrpc-server/jlog"
)

func init() {
	jlog.Init(os.Stdout, "", jlog.LstdFlags|jlog.Lshortfile)
	jlog.SetLevel(jlog.DEBUG)
}

func main() {
	jruntime.Register()
	jlog.Info("jrpc server start")
	err := jRPC.Start(":4241")
	if err != nil {
		fmt.Println(err)
	}
}
