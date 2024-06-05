package protocol

import (
	"time"
)

type EventStatus string

const (
	EventUnhandledStatus = EventStatus("unhandled") // Default event status.
	EventConfirmedStatus = EventStatus("confirmed") // Is event has been confirmed?
	EventAlarmedStatus   = EventStatus("alarmed")   // Is event occurred during alarm?
	EventDeclinedStatus  = EventStatus("declined")  // Is event has been declined?
)

type Event struct {
	ID       string        `json:"id"`
	Start    time.Time     `json:"start"`
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
	Status   EventStatus   `json:"status"`
}

type EventStartMessage struct {
	ID     string      `json:"id"`
	Start  time.Time   `json:"start"`
	Status EventStatus `json:"status"`
}

type EventEndMessage struct {
	ID       string        `json:"id"`
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
	Status   EventStatus   `json:"status"`
}

func NewEventStartMessage(producer Producer, message EventStartMessage) Message[EventStartMessage] {
	return Message[EventStartMessage]{
		Producer: producer,
		Type:     MessageEventStart,
		Body:     message,
	}
}

func NewEventEndMessage(producer Producer, message EventEndMessage) Message[EventEndMessage] {
	return Message[EventEndMessage]{
		Producer: producer,
		Type:     MessageEventEnd,
		Body:     message,
	}
}
