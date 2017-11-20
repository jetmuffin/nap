package console

// Message defines messages communicate with hterm.
type Message struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// ResizeContent defines resize command inside console message.
type ResizeContent struct {
	Columns int `json:"columns"`
	Rows    int `json:"rows"`
}

const (
	// PingMessage is the ping message type sent from hterm.
	PingMessage = iota
	// PongMessage is the pong message type sent to hterm.
	PongMessage
	// InitMessage is the initialize message type sent from hterm.
	InitMessage
	// ResizeMessage is the resize message type sent from hterm.
	ResizeMessage
	// InputMessage is the input message type sent from hterm.
	InputMessage
	// OutputMessage is the output message type sent to hterm.
	OutputMessage
	// ErrorMessage is the error message type sent to hterm.
	ErrorMessage
)
