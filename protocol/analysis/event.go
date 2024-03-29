package analysis

import (
	"time"

	. "github.com/nenavizhuleto/horizon/protocol/detection"
)

type Trigger Detection

type Event struct {
	ID      string
	Trigger Trigger

	Start    time.Time
	End      time.Time
	Duration time.Duration
}
