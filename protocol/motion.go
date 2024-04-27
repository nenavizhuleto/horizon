package protocol

import "time"

type Motion struct {
	// Video stream or file
	Source Source `json:"source"`

	// Position where motion is detected
	Position Position `json:"position,omitempty"`
}

type MotionDetectionMessage = DetectionMessage[Motion]

func NewMotionDetectionMessage(producer Producer, ts time.Time, motions []Motion, options ...MessageOptions) MotionDetectionMessage {
	return NewDetectionMessage(MessageMotionDetection, producer, ts, motions, options...)
}
