package analysis

import (
	. "github.com/nenavizhuleto/horizon/protocol/detection"
)

type Analysis struct {
	Event     Event     `json:"event"`
	Detection Detection `json:"detection"`
	Report    Report    `json:"report"`
}
