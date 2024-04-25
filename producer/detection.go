package producer

import (
	"time"
)

type DetectionProducer[D, M, O any] interface {
	NewDetectionMessage(ts time.Time, options O, detections ...D) M
}
