package errors

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	SUCCESS             = 200
	SERVERINTERNALERROR = 201
	FUNCTIONNOTFOUND    = 202
)

type ErrJrpc struct {
	Code     int
	Msg      string
	fileName string
	lineNum  int
	Err      error
}

func NewErr(code int, msg string, err error) *ErrJrpc {
	ej := new(ErrJrpc)
	_, f, l, ok := runtime.Caller(1)
	if !ok {
		f = "unknown"
		l = 0
	}
	ej.Code = code
	ej.Msg = msg
	ej.Err = err
	ej.fileName = f
	ej.lineNum = l
	return ej
}

func (ej *ErrJrpc) Error() string {
	ej.fileName = ej.fileName[strings.LastIndex(ej.fileName, "/")+1:]
	return fmt.Sprintf("%s:%d: %s\n", ej.fileName, ej.lineNum, ej.Err.Error())
}

func GetMsg(code int) string {
	var s string
	switch code {
	case 200:
		s = "success"
	case 201:
		s = "server internal error"
	case 202:
		s = "function not found"
	}
	return s
}
