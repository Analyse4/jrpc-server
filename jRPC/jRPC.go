package jRPC

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/Analyse4/jrpc-server/jRPC/errors"
	"github.com/Analyse4/jrpc-server/jRPC/stub"
	"github.com/Analyse4/jrpc-server/jlog"
	"github.com/Analyse4/jrpc-server/protocol"
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

			ack := new(protocol.BaseMsgAck)
			err = stub.Handle(buffer[:n], ack)
			if err != nil {
				ack.Code = err.(*errors.ErrJrpc).Code
				ack.Msg = err.(*errors.ErrJrpc).Msg
			} else {
				ack.Code = errors.SUCCESS
				ack.Msg = errors.GetMsg(errors.SUCCESS)
			}
			jack, err := json.Marshal(ack)
			if err != nil {
				err = errors.NewErr(errors.SERVERINTERNALERROR, errors.GetMsg(errors.SERVERINTERNALERROR), fmt.Errorf("json unmarshal error"))
				doneChan <- err
				return
			}
			n, err = pc.WriteTo(jack, caddr)
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
