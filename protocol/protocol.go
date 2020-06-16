package protocol

// BaseMsgReq is basic req msg struct for jRPC
type BaseMsgReq struct {
	ID   string `json:"id"`
	Data []byte `json:"msg"`
}

// BaseMsgAck is basic ack msg struct for jRPC
type BaseMsgAck struct {
	Code int
	Msg  string
	ID   string `json:"id"`
	Data []byte `json:"msg"`
}

type SimpleReq struct {
	Content string `json:"content"`
}

type SimpleAck struct {
	Content string `json:"content"`
}
