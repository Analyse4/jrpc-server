package protocol

// BaseMsg is basic msg struct for jRPC
type BaseMsg struct {
	ID  string `json:"id"`
	Msg []byte `json:"msg"`
}

type SimpleReq struct {
	Content string `json:"content"`
}

type SimpleAck struct {
	Content string `json:"content"`
}
