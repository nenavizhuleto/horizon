package protocol

import "encoding/json"

type MessageType string

const (
	MessageMotionDetection = MessageType("detection.motion")
	MessageObjectDetection = MessageType("detection.object")
	MessagePlateDetection  = MessageType("detection.plate")

	MessageFrame = MessageType("frame")
	MessageMedia = MessageType("media")

	MessageEventStart = MessageType("event.start")
	MessageEventEnd   = MessageType("event.end")

	MessageAnalysis       = MessageType("analysis")
	MessageObjectAnalysis = MessageType("analysis.object")
	MessagePlateAnalysis  = MessageType("analysis.plate")
)

type Message[B any] struct {
	Producer Producer    `json:"producer"`
	Type     MessageType `json:"type"`
	Body     B           `json:"body"`
}

type RawMessage Message[json.RawMessage]
type RawAnalysesMessage AnalysesMessage[json.RawMessage]
type RawDetectionMessage DetectionMessage[json.RawMessage]
