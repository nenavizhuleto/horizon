package protocol

import "time"

type Detection interface{}

type DetectionMessage[D any] struct {
	ProducerMessage[Producer]
	Timestamp  time.Time `json:"timestamp"`
	Detections []D       `json:"detections"`
}

func NewDetectionMessage[D any](t MessageType, producer Producer, ts time.Time, detections []D, options ...MessageOptions) DetectionMessage[D] {
	return DetectionMessage[D]{
		ProducerMessage: NewProducerMessage(t, producer, options...),
		Timestamp:       ts,
		Detections:      detections,
	}
}
