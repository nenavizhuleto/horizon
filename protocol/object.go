package protocol

import "time"

type Object struct {
	// Arbitrary value
	//
	// Example: person, car, ball, bycicle
	Class string `json:"class"`

	// Where object was detected
	BoundingBox Position `json:"bounding_box"`

	// How confident this detection is
	//
	// Floating-point number: 0.0 <= x <= 1.0
	Confidence float32 `json:"confidence"`

	// We also might want to have
	// additional information alongise with detection
	UserData any `json:"user_data"`
}

type ObjectDetectionMessage = DetectionMessage[Object]

func NewObjectDetectionMessage(producer Producer, ts time.Time, objects []Object, options ...MessageOptions) ObjectDetectionMessage {
	return NewDetectionMessage(MessageObjectDetection, producer, ts, objects, options...)
}
