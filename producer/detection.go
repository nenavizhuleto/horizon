package producer

import (
	"time"
)

type DetectionProducer[D, M any] interface {
	NewDetectionMessage(ts time.Time, detections ...D) M
}

type DetectionExtProducer[D, M, O any] interface {
	NewExtDetectionMessage(ts time.Time, options O, detections ...D) M
}
