package protocol

type Frame []byte

type FrameMessageOptions struct {
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	Regions    *[]Position `json:"regions,omitempty"`
}

type FrameMessage struct {
	FrameMessageOptions
	ProducerMessage[Producer]
	Frame Frame `json:"frame"`
}

func NewFrameMessage(producer Producer, frame Frame, options FrameMessageOptions) FrameMessage {
	return FrameMessage{
		ProducerMessage:     NewProducerMessage(MessageFrame, producer),
		Frame:               frame,
		FrameMessageOptions: options,
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
