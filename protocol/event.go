package protocol

import (
	"time"
)

type Event struct {
	ID       string        `json:"id"`
	Trigger  Detection     `json:"trigger"`
	Start    time.Time     `json:"start"`
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
}

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

func NewEventMessage(t MessageType, producer Producer, id string, options ...MessageOptions) EventMessage {
	return EventMessage{
		ProducerMessage: NewProducerMessage(t, producer, options...),
		ID:              id,
	}
}

func NewEventStartMessage(producer Producer, id string, start time.Time, trigger Detection, options ...MessageOptions) EventStartMessage {
	return EventStartMessage{
		EventMessage: NewEventMessage(MessageEventStart, producer, id, options...),
		Start:        start,
		Trigger:      trigger,
	}
}

func NewEventEndMessage(producer Producer, id string, end time.Time, duration time.Duration, options ...MessageOptions) EventEndMessage {
	return EventEndMessage{
		EventMessage: NewEventMessage(MessageEventEnd, producer, id, options...),
		End:          end,
		Duration:     duration,
	}
}
