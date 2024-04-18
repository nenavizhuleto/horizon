package protocol

import (
	"time"
)

type DetectionType string

// Primitive detection types
const (
	// Detect motion on video
	MotionDetection = DetectionType("motion")

	// Detect object on image/frame
	ObjectDetection = DetectionType("object")

	// Detect value on analog/digital sensor input
	ValueDetection = DetectionType("value")
)

type Detection struct {
	// Which sensor made detection
	Producer Sensor `json:"producer"`

	// Although each sensor has it's own type
	// and associated detection types with it,
	// it would be nice to pass detection type further,
	// for better API
	Type DetectionType `json:"detection_type"`

	// When did detection happened
	Timestamp time.Time `json:"timestamp"`

	// As detection value may vary based on different sensor types,
	// We should be able to pass any data to consumers.
	// So also we need some way to extract this value in sensor intended way.
	Value any `json:"value"`
}
