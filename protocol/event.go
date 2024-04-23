package protocol

import (
	"time"
)

type Trigger Detection

type Event struct {
	ID      string  `json:"id"`
	Trigger Trigger `json:"trigger"`

	Start    time.Time     `json:"start"`
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
}

func NewEvent(id string, trigger Trigger) Event {
	return Event{
		ID:       id,
		Trigger:  trigger,
		Start:    time.Time{},
		End:      time.Time{},
		Duration: 0,
	}
}

func (e *Event) StartNow() {
	e.Start = time.Now()
}

func (e *Event) EndNow() {
	e.End = time.Now()
	e.Duration = e.End.Sub(e.Start)
}
