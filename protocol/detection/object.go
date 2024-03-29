package detection

import "github.com/nenavizhuleto/horizon/protocol/detection/video"

type Object struct {
	// Arbitrary value
	//
	// Example: person, car, ball, bycicle
	Class string

	// Where object was detected
	BoundingBox video.Position

	// How confident this detection is
	//
	// Floating-point number: 0.0 <= x <= 1.0
	Confidence float32
}
