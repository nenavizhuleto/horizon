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
