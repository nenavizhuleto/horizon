package protocol

import (
	"time"
)

type EventMessage struct {
	ProducerMessage[Producer]
	ID string `json:"id"`
}

type EventStartMessage struct {
	EventMessage
	Trigger Detection `json:"trigger"`
	Start   time.Time `json:"start"`
}

type EventEndMessage struct {
	EventMessage
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
}
