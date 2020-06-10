package main

import (
	"fmt"

	"github.com/Analyse4/jrpc-server/jRPC"
)

func main() {
	fmt.Println("server start")
	jRPC.Start(":4242")
}
