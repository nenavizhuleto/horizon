package protocol

type MessageType string

const (
	MessageDetection = MessageType("detection")

	MessageMotion = MessageType("motion")
	MessageObject = MessageType("object")
	MessagePlate  = MessageType("plate")

	MessageFrame = MessageType("frame")
	MessageMedia = MessageType("media")

	MessageEvent = MessageType("event")
	MessageStart = MessageType("start")
	MessageEnd   = MessageType("end")

	MessageEventStart = MessageType("event.start")
	MessageEventEnd   = MessageType("event.end")

	MessageAnalysis       = MessageType("analysis")
	MessageObjectAnalysis = MessageType("analysis.object")
	MessagePlateAnalysis  = MessageType("analysis.plate")
)

const MessageTypeDelimiter = "."

func NewMessageType(base MessageType, parts ...MessageType) MessageType {
	for _, p := range parts {
		base = MessageType(string(base) + MessageTypeDelimiter + string(p))
	}
	return base
}

type Body interface {
	DetectionMessage[PlateDetection] | DetectionMessage[ObjectDetection] | AnalysesMessage[any, PlateDetection] | AnalysesMessage[any, ObjectDetection]
	Type() MessageType
}

type Message[
	P Producer,
	B Body,
] struct {
	Producer P           `json:"producer"`
	Type     MessageType `json:"type"`
	Body     B           `json:"body"`
}

type Producer interface {
	Camera | Service
}

func NewMessage[
	P Camera | Service,
	B Body,
](producer P, body B) Message[P, B] {
	return Message[P, B]{
		Producer: producer,
		Type:     body.Type(),
		Body:     body,
	}
}

func FOO() {
	a := NewAnalysesMessage(PlateDetection{}, "")
	m1 := NewMessage(Service{}, a)

	m2 := NewMessage(Camera{}, AnalysesMessage[any, PlateDetection]{
		Test: PlateDetection{},
	})
}
