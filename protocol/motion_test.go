package protocol_test

import (
	"time"

	. "github.com/nenavizhuleto/horizon/protocol"
	"github.com/nenavizhuleto/horizon/protocol/video"
)

func TestMotionDetection() {
	_ = Detection{
		Producer: Sensor{
			ID:   "xxxx-xxxx-xxxx-xxxx",
			Name: "my-beloved-camera",
			Type: VideoSensor,
		},
		Type:      MotionDetection,
		Timestamp: time.Now(),
		Value: Motion{
			Source: video.Source{
				URI: "rtsp://localhost:8554/camera",
				Dimensions: video.Dimensions{
					Width:  1920,
					Height: 1080,
				},
			},
			Position: &video.Position{
				X:      200,
				Y:      400,
				Width:  500,
				Height: 300,
			},
		},
	}
}
