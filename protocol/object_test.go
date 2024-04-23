package protocol_test

import (
	"testing"
	"time"

	. "github.com/nenavizhuleto/horizon/protocol"
	"github.com/nenavizhuleto/horizon/protocol/video"
)

func TestObjectDetection(t *testing.T) {
	_ = Detection{
		Producer: Sensor{
			ID:   "xxxx-xxxx-xxxx-xxxx",
			Name: "license-plate-recognition",
			Type: ObjectSensor,
		},
		Type:      ObjectDetection,
		Timestamp: time.Now(),
		Value: Object{
			Class: "license plate",
			BoundingBox: video.Position{
				X:      300,
				Y:      200,
				Width:  300,
				Height: 300,
			},
			Confidence: 0.98,
		},
	}
}
