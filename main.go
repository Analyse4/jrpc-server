package main

import (
	"fmt"
	"os"

	"github.com/Analyse4/jrpc-server/jRPC"
	jruntime "github.com/Analyse4/jrpc-server/jRPC/runtime"
	"github.com/Analyse4/jrpc-server/jlog"
)

var jl *jlog.JLogger

func init() {
	jl = jlog.New(os.Stdout, "", jlog.LstdFlags|jlog.Lshortfile)
	//jl = jlog.New(os.Stdout, "", 0)
	jl.SetLevel(jlog.INFO)
}

func main() {
	jl.Info("test1")
	jl.Debug("test2")

	fmt.Println("server start")
	jruntime.Register()
	err := jRPC.Start(":4241")
	if err != nil {
		fmt.Println(err)
	}
}
