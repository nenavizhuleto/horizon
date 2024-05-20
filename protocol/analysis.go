package protocol

import "time"

type Severity string

const (
	SeverityInfo    = Severity("info")
	SeverityWarning = Severity("warning")
	SeverityPanic   = Severity("panic")
)

type Analysis[R, S any] struct {
	ID      string `json:"id"`
	Report  R      `json:"report"`
	Subject S      `json:"subject"`
}

func NewAnalysis[R, S any](id string, report R, subject S) Analysis[R, S] {
	return Analysis[R, S]{
		ID:      id,
		Report:  report,
		Subject: subject,
	}
}

type AnalysisMessage[A any] struct {
	EventID   string    `json:"event_id"`
	Timestamp time.Time `json:"timestamp"`
	Severity  Severity  `json:"severity"`
	Location  Location  `json:"location"`
	Analyses  []A       `json:"analyses"`
}
