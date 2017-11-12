package console

type ConsoleMessage struct {
	Type int `json:"type"`
	Content string `json:"content"`
}

type ResizeContent struct {
	Columns int `json:"columns"`
	Rows int `json:"rows"`
}

const (
	PingMessage   = iota
	PongMessage
	InitMessage
	ResizeMessage
	InputMessage
	OutputMessage
	ErrorMessage
)
