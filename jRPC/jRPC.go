package jRPC

import (
	"fmt"
	"net"
	"time"

	"github.com/Analyse4/jrpc-server/jRPC/stub"
)

const maxBufferSize = 1024

// Start wrap with stub.send
func Start(addr string) error {
	pc, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	defer pc.Close()
	doneChan := make(chan error, 1)
	buffer := make([]byte, maxBufferSize)

	go func() {
		for {
			n, caddr, err := pc.ReadFrom(buffer)
			if err != nil {
				doneChan <- err
				return
			}
			fmt.Printf("%v  datagram received: bytes: %d, from: %s\n", time.Now(), n, caddr.String())
			ack, err := stub.Handle(buffer)
			if err != nil {
				doneChan <- err
				return
			}
			n, err = pc.WriteTo(ack, caddr)
			if err != nil {
				doneChan <- err
				return
			}
			fmt.Printf("%v  datagram writed: byets: %d, to: %s\n", time.Now(), n, caddr.String())
		}
	}()

	select {
	case err = <-doneChan:
	}
	// TODO: Maybe need inform client
	return err
}
