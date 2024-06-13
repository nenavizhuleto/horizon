package protocol

import (
	"time"
)

type Severity string

const (
	SeverityInfo    = Severity("info")
	SeverityWarning = Severity("warning")
	SeverityPanic   = Severity("panic")
)

type Analysis[R any, S Typable] struct {
	ID      string `json:"id"`
	Report  R      `json:"report"`
	Subject S      `json:"subject"`
}

func NewAnalysis[R any, S Typable](id string, report R, subject S) Analysis[R, S] {
	return Analysis[R, S]{
		ID:      id,
		Report:  report,
		Subject: subject,
	}
}

type AnalysesMessageBody[R any, S Typable] struct {
	EventID       string           `json:"event_id"`
	Timestamp     time.Time        `json:"timestamp"`
	Severity      Severity         `json:"severity"`
	FrameLocation FrameLocation    `json:"frame_location"`
	Analyses      []Analysis[R, S] `json:"analyses"`
}

func (m AnalysesMessageBody[R, S]) Type() MessageType {
	var s S
	return Join(MessageAnalysis, s.Type())
}
