package protocol

import "time"

type Frame []byte

type FrameMessage struct {
	Timestamp  time.Time  `json:"timestamp"`
	Dimensions Dimensions `json:"dimensions,omitempty"`
	Regions    []Position `json:"regions,omitempty"`
	Data       Frame      `json:"data"`
}

func NewFrameMessage(producer Camera, message FrameMessage) Message[FrameMessage] {
	return Message[FrameMessage]{
		Producer: producer,
		Type:     MessageFrame,
		Body:     message,
	}
}

type Source struct {
	// Uniform Resource Identifier
	URI string `json:"uri"`

	// Width(X) and Height(Y) of the source, on which consumers could rely on
	Dimensions Dimensions `json:"dimensions"`
}

// Might used image.Point, but implemented for better marshal control
type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Might used image.Rectangle, but will not
type Position struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}
