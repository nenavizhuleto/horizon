package protocol

import (
	"encoding/json"
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

func (d *Detection) UnmarshalJSON(data []byte) error {
	type DTO struct {
		Producer  Sensor          `json:"producer"`
		Type      DetectionType   `json:"detection_type"`
		Timestamp time.Time       `json:"timestamp"`
		Value     json.RawMessage `json:"value"`
	}

	var dto DTO
	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	var (
		value any
		err   error
	)

	switch dto.Type {
	case ObjectDetection:
		value, err = extractor.Raw[Object](dto.Value)
	case MotionDetection:
		value, err = extractor.Raw[Motion](dto.Value)
	case ValueDetection:
		value, err = extractor.Raw[Value](dto.Value)
	}

	if err != nil {
		return err
	}

	d.Producer = dto.Producer
	d.Type = dto.Type
	d.Timestamp = dto.Timestamp
	d.Value = value

	return nil
}
