package detection

import "github.com/nenavizhuleto/horizon/protocol/detection/video"

type Motion struct {
	// Video stream or file
	Source video.Source

	// Position where motion is detected (optional)
	Position *video.Position
}
