package stub

import (
	"encoding/json"
	"fmt"

	"github.com/Analyse4/jrpc-server/jRPC/errors"
	jruntime "github.com/Analyse4/jrpc-server/jRPC/runtime"
	"github.com/Analyse4/jrpc-server/protocol"
)

// Handle return json format response
func Handle(msg []byte, bma *protocol.BaseMsgAck) error {
	//msg = bytes.Trim(msg, "\x00")

	bm := new(protocol.BaseMsgReq)
	bm.Data = make([]byte, 0)

	err := json.Unmarshal(msg, bm)
	if err != nil {
		return errors.NewErr(errors.SERVERINTERNALERROR, errors.GetMsg(errors.SERVERINTERNALERROR), fmt.Errorf("json unmarshal error"))
	}

	data, err := jruntime.Handle(bm)
	if err != nil {
		return err
	}

	bma.ID = bm.ID
	bma.Data = data
	// jbm, err := json.Marshal(bm)
	// if err != nil {
	// 	return []byte(""), errors.NewErr(errors.SERVERINTERNALERROR, errors.GetMsg(errors.SERVERINTERNALERROR), fmt.Errorf("json marshal error"))
	// }
	return nil
}
