package jRPC

import (
	"net"

	"github.com/Analyse4/jrpc-server/jRPC/stub"
	"github.com/Analyse4/jrpc-server/jlog"
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
			jlog.Debugf("datagram received: bytes: %d, from: %s\n", n, caddr.String())

			ack, err := stub.Handle(buffer[:n])
			if err != nil {
				doneChan <- err
				return
			}
			n, err = pc.WriteTo(ack, caddr)
			if err != nil {
				doneChan <- err
				return
			}
			jlog.Debugf("datagram writed: byets: %d, to: %s\n", n, caddr.String())
		}
	}()

	select {
	case err = <-doneChan:
	}
	// TODO: Maybe need inform client
	return err
}
