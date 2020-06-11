package main

import (
	"fmt"

	"github.com/Analyse4/jrpc-server/jRPC"
	jruntime "github.com/Analyse4/jrpc-server/jRPC/runtime"
)

func main() {
	fmt.Println("server start")
	jruntime.Register()
	err := jRPC.Start(":4241")
	if err != nil {
		fmt.Println(err)
	}
}
