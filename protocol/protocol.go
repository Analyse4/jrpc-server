package protocol

// BaseMsg is basic msg struct for jRPC
type BaseMsg struct {
	ID  string
	Msg []byte
}
