package protocol

import "time"

type MotionDetection struct {
	Source   Source   `json:"source"`
	Position Position `json:"position,omitempty"`
}

type MotionDetectionMessage struct {
	Timestamp  time.Time         `json:"timestamp"`
	Detections []MotionDetection `json:"detections"`
}

func NewMotionDetectionMessage(producer Camera, message MotionDetectionMessage) Message[MotionDetectionMessage] {
	return Message[MotionDetectionMessage]{
		Producer: producer,
		Type:     MessageMotionDetection,
		Body:     message,
	}
}
