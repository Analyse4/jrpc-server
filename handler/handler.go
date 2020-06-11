package handler

import (
	"encoding/json"

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
		return data, err
	}
	return data, nil
}
