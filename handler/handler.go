package handler

import (
	"encoding/json"
	"fmt"

	"github.com/Analyse4/jrpc-server/jRPC/errors"

	"github.com/Analyse4/jrpc-server/protocol"
)

// Handler is a struct for register conveniently
type Handler struct{}

// SimpleHandler is a simple handler for test
func (h *Handler) SimpleHandler(req *protocol.SimpleReq) ([]byte, error) {
	ack := new(protocol.SimpleAck)
	ack.Content = "ack"
	data, err := json.Marshal(ack)
	if err != nil {
		return data, errors.NewErr(errors.SERVERINTERNALERROR, errors.GetMsg(errors.SERVERINTERNALERROR), fmt.Errorf("json marshal error"))
	}
	return data, nil
	//return data, errors.NewErr(500, "server internal error", fmt.Errorf("test-errir"))
}
