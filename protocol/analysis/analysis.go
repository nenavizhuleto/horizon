package analysis

import (
	. "github.com/nenavizhuleto/horizon/protocol/detection"
)

type Analysis struct {
	Event     Event
	Detection Detection
	Report    Report
}
