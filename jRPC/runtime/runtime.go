package runtime

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/Analyse4/jrpc-server/handler"
	"github.com/Analyse4/jrpc-server/jlog"
	"github.com/Analyse4/jrpc-server/protocol"
)

var jl *jlog.JLogger
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
	var err error
	jl, err = jlog.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
	jl.Debug(strings.Join(logm, ","))
}

// Handle return the result data by invoking method which is finded by rpc ID
func Handle(msg *protocol.BaseMsg) ([]byte, error) {
	data := make([]byte, 1024)
	hm, ok := hMap[msg.ID]
	if !ok {
		return data, fmt.Errorf("coressponding handler is not exist")
	}
	req := reflect.New(hm.input.Elem())
	err := json.Unmarshal(msg.Msg, req.Interface())
	if err != nil {
		return data, err
	}
	result := hm.function.Func.Call([]reflect.Value{reflect.ValueOf(new(handler.Handler)), req})
	if result[1].Interface() != nil {
		return data, fmt.Errorf(result[1].String())
	}
	return result[0].Bytes(), nil
}
