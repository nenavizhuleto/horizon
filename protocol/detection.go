package protocol

import "time"

type Detection interface{}

type DetectionMessage[T any] struct {
	ProducerMessage[T]
	Timestamp  time.Time   `json:"timestamp"`
	Detections []Detection `json:"detections"`
}

func NewDetectionMessage[T any](t MessageType, producer T, ts time.Time) DetectionMessage[T] {
	return DetectionMessage[T]{
		ProducerMessage: NewProducerMessage(t, producer),
		Timestamp:       ts,
	}
}
