package protocol_test

import (
	"time"

	. "github.com/nenavizhuleto/horizon/protocol"
	"github.com/nenavizhuleto/horizon/protocol/video"
)

func TestObjectDetection() {
	_ = Detection{
		Producer: Sensor{
			ID:   "xxxx-xxxx-xxxx-xxxx",
			Name: "license-plate-recognition",
			Type: NeuralSensor,
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
