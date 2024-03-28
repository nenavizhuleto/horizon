package detection_test

import (
	"time"

	. "github.com/nenavizhuleto/horizon/protocol/detection"
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
			Source: VideoSource{
				URI: "rtsp://localhost:8554/camera",
				Dimensions: VideoDimensions{
					Width:  1920,
					Height: 1080,
				},
			},
			Position: &VideoPosition{
				X:      200,
				Y:      400,
				Width:  500,
				Height: 300,
			},
		},
	}
}
