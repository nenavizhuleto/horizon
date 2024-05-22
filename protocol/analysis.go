package protocol

import "time"

type Severity string

const (
	SeverityInfo    = Severity("info")
	SeverityWarning = Severity("warning")
	SeverityPanic   = Severity("panic")
)

type Analysis[R, S any] struct {
	Report  R `json:"report"`
	Subject S `json:"subject"`
}

func NewAnalysis[R, S any](report R, subject S) Analysis[R, S] {
	return Analysis[R, S]{
		Report:  report,
		Subject: subject,
	}
}

type AnalysesMessage[A any] struct {
	ID            string        `json:"id"`
	EventID       string        `json:"event_id"`
	Timestamp     time.Time     `json:"timestamp"`
	Severity      Severity      `json:"severity"`
	FrameLocation FrameLocation `json:"frame_location"`
	Analyses      []A           `json:"analyses"`
}
