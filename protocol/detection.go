package protocol

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nenavizhuleto/horizon/protocol/extractor"
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

func (d Detection) IsMotion() bool {
	return d.Type == MotionDetection
}

func (d Detection) Motions() (motions []Motion) {
	if d.IsMotion() {
		motions = d.Value.([]Motion)
	}

	return
}

func (d Detection) IsObject() bool {
	return d.Type == ObjectDetection
}

func (d Detection) Objects() (objects []Object) {
	if d.IsObject() {
		objects = d.Value.([]Object)
	}

	return
}

func (d Detection) IsValue() bool {
	return d.Type == ValueDetection
}

func (d Detection) Values() (values []Value) {
	if d.IsValue() {
		values = d.Value.([]Value)
	}

	return
}

func (d *Detection) SetTimestamp(ts time.Time) {
	d.Timestamp = ts
}

func (d *Detection) Range(foreach func(i int, v any) bool) {
	values, ok := d.Value.([]any)
	if ok {
		for i, value := range values {
			if !foreach(i, value) {
				break
			}
		}
	}
}

func (d *Detection) UnmarshalJSON(data []byte) error {
	type DetectionJSON Detection

	valueJSON := &json.RawMessage{}
	detectionJSON := DetectionJSON{
		Value: valueJSON,
	}

	if err := json.Unmarshal(data, &detectionJSON); err != nil {
		return err
	}

	var (
		value any
		err   error
	)

	switch detectionJSON.Type {
	case ObjectDetection:
		value, err = extractor.Raw[[]Object](*valueJSON)
	case MotionDetection:
		value, err = extractor.Raw[[]Motion](*valueJSON)
	case ValueDetection:
		value, err = extractor.Raw[[]Value](*valueJSON)
	default:
		return fmt.Errorf("unsupported detection type: %s", detectionJSON.Type)
	}

	if err != nil {
		return err
	}

	*d = Detection(detectionJSON)
	d.Value = value

	return nil
}
