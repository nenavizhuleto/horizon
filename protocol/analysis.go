package protocol

import "time"

type Severity string

const (
	SeverityInfo    = Severity("info")
	SeverityWarning = Severity("warning")
	SeverityPanic   = Severity("panic")
)

type Analysis[D any] struct {
	Report    any `json:"report"`
	Candidate D   `json:"candidate"`
}

type MotionAnalysis = Analysis[Motion]
type ObjectAnalysis = Analysis[Object]
type ValueAnalysis = Analysis[Value]

func NewAnalysis[D any](candidate D, report any) Analysis[D] {
	return Analysis[D]{
		Report:    report,
		Candidate: candidate,
	}
}

type AnalysisMessage[C any] struct {
	ProducerMessage[Producer]
	Timestamp time.Time     `json:"timestamp"`
	Severity  Severity      `json:"severity"`
	Analyses  []Analysis[C] `json:"analyses"`
}

func NewAnalysisMessage[C any](producer Producer, ts time.Time, severity Severity, analyses []Analysis[C], options ...MessageOptions) AnalysisMessage[C] {
	return AnalysisMessage[C]{
		ProducerMessage: NewProducerMessage(MessageAnalysis, producer, options...),
		Timestamp:       ts,
		Severity:        severity,
		Analyses:        analyses,
	}
}
