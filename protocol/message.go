package protocol

type MessageType string

const (
	MessageMotionDetection    = MessageType("detection.motion")
	MessageObjectDetection    = MessageType("detection.object")
	MessageObjectDetectionExt = MessageType("detection.object.ext")
	MessageValueDetection     = MessageType("detection.value")

	MessageFrame    = MessageType("frame")
	MessageFrameExt = MessageType("frame.ext")

	MessageEvent      = MessageType("event")
	MessageEventStart = MessageType("event.start")
	MessageEventEnd   = MessageType("event.end")

	MessageAnalysisDetection = MessageType("analysis.detection")
	MessageAnalysisEvent     = MessageType("analysis.event")
)

type MessageHeaders map[string]any

type Message struct {
	Type MessageType `json:"type"`
}

func new_message(t MessageType) Message {
	return Message{
		Type: t,
	}
}
