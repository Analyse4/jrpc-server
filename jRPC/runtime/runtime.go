package runtime

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/Analyse4/jrpc-server/handler"
	"github.com/Analyse4/jrpc-server/jRPC/errors"
	"github.com/Analyse4/jrpc-server/jlog"
	"github.com/Analyse4/jrpc-server/protocol"
)

var hMap map[string]*handlerMeta

type handlerMeta struct {
	function reflect.Method
	input    reflect.Type
}

func init() {
	hMap = make(map[string]*handlerMeta)
}

// Register is used to register handler
func Register() {
	var logm []string
	h := new(handler.Handler)
	hTyp := reflect.TypeOf(h)
	for i := 0; i < hTyp.NumMethod(); i++ {
		hm := new(handlerMeta)
		hm.function = hTyp.Method(i)
		hm.input = hm.function.Type.In(1)

		id := "jrpc." + strings.ToLower(hm.function.Name)
		hMap[id] = hm

		logm = append(logm, id)
	}
	jlog.Debug(strings.Join(logm, ","))
}

// Handle return the result data by invoking method which is finded by rpc ID
func Handle(msg *protocol.BaseMsgReq) ([]byte, error) {
	data := make([]byte, 1024)
	hm, ok := hMap[msg.ID]
	if !ok {
		return data, errors.NewErr(errors.FUNCTIONNOTFOUND, errors.GetMsg(errors.FUNCTIONNOTFOUND), fmt.Errorf("function %s isn't registed", msg.ID))
	}
	req := reflect.New(hm.input.Elem())
	err := json.Unmarshal(msg.Data, req.Interface())
	if err != nil {
		return data, errors.NewErr(errors.SERVERINTERNALERROR, errors.GetMsg(errors.SERVERINTERNALERROR), fmt.Errorf("json unmarshal error"))
	}
	result := hm.function.Func.Call([]reflect.Value{reflect.ValueOf(new(handler.Handler)), req})
	if result[1].Interface() != nil {
		return data, result[1].Interface().(*errors.ErrJrpc)
	}
	return result[0].Bytes(), nil
}
