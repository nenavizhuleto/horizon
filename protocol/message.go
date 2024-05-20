package protocol

type MessageType string

const (
	MessageMotionDetection = MessageType("detection.motion")

	MessageObjectDetection = MessageType("detection.object")
	MessagePlateDetection  = MessageType("detection.plate")
	MessagePersonDetection = MessageType("detection.person")

	MessageFrame = MessageType("frame")
	MessageMedia = MessageType("media")

	MessageEventStart = MessageType("event.start")
	MessageEventEnd   = MessageType("event.end")

	MessageObjectAnalysis = MessageType("analysis.object")
	MessagePlateAnalysis  = MessageType("analysis.plate")
	MessagePersonAnalysis = MessageType("analysis.person")
)

type Message[B any] struct {
	Producer Producer    `json:"producer"`
	Type     MessageType `json:"type"`
	Body     B           `json:"body"`
}
