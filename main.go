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
	jl.SetLevel(jlog.DEBUG)
}

func main() {
	jl.Info("jrpc server start")
	jruntime.Register()
	err := jRPC.Start(":4241")
	if err != nil {
		fmt.Println(err)
	}
}
