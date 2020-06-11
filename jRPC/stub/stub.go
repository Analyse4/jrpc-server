package stub

import (
	"bytes"
	"encoding/json"
	"fmt"

	jruntime "github.com/Analyse4/jrpc-server/jRPC/runtime"
	"github.com/Analyse4/jrpc-server/protocol"
)

// Handle return json format response
func Handle(msg []byte) ([]byte, error) {
	msg = bytes.Trim(msg, "\x00")

	bm := new(protocol.BaseMsg)
	bm.Msg = make([]byte, 1024)

	fmt.Println(string(msg))
	err := json.Unmarshal(msg, bm)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	data, err := jruntime.Handle(bm)
	if err != nil {
		return []byte(""), err
	}

	bm.Msg = data
	jbm, err := json.Marshal(bm)
	if err != nil {
		return []byte(""), err
	}
	return jbm, nil
}
