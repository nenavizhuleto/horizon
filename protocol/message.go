package protocol

type MessageType string

const (
	MessageMotionDetection = MessageType("detection.motion")
	MessageObjectDetection = MessageType("detection.object")
	MessageValueDetection  = MessageType("detection.value")

	MessageFrame = MessageType("frame")

	MessageEvent      = MessageType("event")
	MessageEventStart = MessageType("event.start")
	MessageEventEnd   = MessageType("event.end")

	MessageAnalysis                = MessageType("analysis")
	MessageAnalysisDetection       = MessageType("analysis.detection")
	MessageAnalysisMotionDetection = MessageType("analysis.detection.motion")
	MessageAnalysisObjectDetection = MessageType("analysis.detection.object")
	MessageAnalysisValueDetection  = MessageType("analysis.detection.value")
	MessageAnalysisEvent           = MessageType("analysis.event")
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
