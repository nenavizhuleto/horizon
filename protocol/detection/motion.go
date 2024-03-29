package detection

import "github.com/nenavizhuleto/horizon/protocol/detection/video"

type Motion struct {
	// Video stream or file
	Source video.Source `json:"source"`

	// Position where motion is detected (optional)
	Position *video.Position `json:"position,omitempty"`
}
