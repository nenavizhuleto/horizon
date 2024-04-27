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
	MessageOptions
	Type MessageType `json:"type"`
}

type Location struct {
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Topic     string `json:"topic"`
}

type FrameOptions struct {
	Location   *Location   `json:"location"`
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	Regions    *[]Position `json:"regions,omitempty"`
}

type EventOptions struct {
	ID *string `json:"id,omitempty"`
}

type MessageOptions struct {
	Frame *FrameOptions `json:"frame,omitempty"`
	Event *EventOptions `json:"event,omitempty"`
}

func NewMessage(t MessageType, options ...MessageOptions) Message {
	return Message{
		Type:           t,
		MessageOptions: Options(options...),
	}
}
