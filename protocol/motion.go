package protocol

import "time"

type MotionDetection struct {
	Source   Source   `json:"source"`
	Position Position `json:"position,omitempty"`
}

type MotionDetectionMessageBody struct {
	Timestamp  time.Time         `json:"timestamp"`
	Detections []MotionDetection `json:"detections"`
}

func (m MotionDetectionMessageBody) Type() MessageType {
	return Join(MessageDetection, "motion")
}
