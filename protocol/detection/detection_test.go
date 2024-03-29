package detection_test

import (
	. "github.com/nenavizhuleto/horizon/protocol/detection"
)

var (
	Camera = Sensor{
		ID:   "xxxx-xxxx-xxxx-xxxx",
		Name: "my-beloved-camera",
		Type: VideoSensor,
	}

	LPR = Sensor{
		ID:   "xxxx-xxxx-xxxx-xxxx",
		Name: "license-plate-recognition",
		Type: NeuralSensor,
	}
)
