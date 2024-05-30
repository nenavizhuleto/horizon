package protocol

import (
	"time"
)

type Event struct {
	ID       string        `json:"id"`
	Start    time.Time     `json:"start"`
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
}

type EventStartMessage struct {
	ID    string    `json:"id"`
	Start time.Time `json:"start"`
}

type EventEndMessage struct {
	ID       string        `json:"id"`
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
}

func NewEventStartMessage(producer Camera, message EventStartMessage) Message[EventStartMessage] {
	return Message[EventStartMessage]{
		Producer: producer,
		Type:     MessageEventStart,
		Body:     message,
	}
}

func NewEventEndMessage(producer Camera, message EventEndMessage) Message[EventEndMessage] {
	return Message[EventEndMessage]{
		Producer: producer,
		Type:     MessageEventEnd,
		Body:     message,
	}
}
