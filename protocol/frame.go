package protocol

type Frame []byte

type FrameMessage struct {
	ProducerMessage[Producer]
	Frame Frame `json:"frame"`
}

type ExtFrameMessage struct {
	FrameMessage
	Dimensions Dimensions `json:"dimensions"`
	Regions    []Position `json:"regions"`
}

func NewFrameMessage(producer Producer, frame Frame) FrameMessage {
	return FrameMessage{
		ProducerMessage: NewProducerMessage(MessageFrame, producer),
		Frame:           frame,
	}
}

func NewExtFrameMessage(producer Producer, frame Frame, dims Dimensions, regs ...Position) ExtFrameMessage {
	return ExtFrameMessage{
		FrameMessage: NewFrameMessage(producer, frame),
		Dimensions:   dims,
		Regions:      regs,
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
