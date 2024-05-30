package protocol

import "time"

type Severity string

const (
	SeverityInfo    = Severity("info")
	SeverityWarning = Severity("warning")
	SeverityPanic   = Severity("panic")
)

type Analysis[R any, S Detection] struct {
	Report  R `json:"report"`
	Subject S `json:"subject"`
}

func (a Analysis[R, S]) Type() MessageType {
	return a.Subject.Type()
}

func NewAnalysis[R any, S Detection](report R, subject S) Analysis[R, S] {
	return Analysis[R, S]{
		Report:  report,
		Subject: subject,
	}
}

type AnalysesMessage[R any, D Detection] struct {
	ID            string           `json:"id"`
	EventID       string           `json:"event_id"`
	Timestamp     time.Time        `json:"timestamp"`
	Severity      Severity         `json:"severity"`
	FrameLocation FrameLocation    `json:"frame_location"`
	Analyses      []Analysis[R, D] `json:"analyses"`
	Test          D
}

func NewAnalysesMessage[R any, D Detection](subject D, report R) AnalysesMessage[R, D] {
	return AnalysesMessage[R, D]{
		Analyses: []Analysis[R, D]{},
	}
}

func (m AnalysesMessage[R, D]) Type() MessageType {
	var d D
	return NewMessageType(MessageAnalysis, d.Type())
}
