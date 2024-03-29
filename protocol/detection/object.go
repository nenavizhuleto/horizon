package detection

import "github.com/nenavizhuleto/horizon/protocol/detection/video"

type Object struct {
	// Arbitrary value
	//
	// Example: person, car, ball, bycicle
	Class string `json:"class"`

	// Where object was detected
	BoundingBox video.Position `json:"bounding_box"`

	// How confident this detection is
	//
	// Floating-point number: 0.0 <= x <= 1.0
	Confidence float32 `json:"confidence"`

	// We also might want to have
	// additional information alongise with detection
	UserData any `json:"user_data"`
}
