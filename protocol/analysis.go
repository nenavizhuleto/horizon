package protocol

type Severity string

const (
	SeverityInfo    = Severity("info")
	SeverityWarning = Severity("warning")
	SeverityPanic   = Severity("panic")
)

type AnalysisMessage struct {
	ProducerMessage[Producer]
	Severity Severity `json:"severity"`
}

type EventAnalysisMessage struct {
	EventMessage
	Detection Detection `json:"detection"`
}

type DetectionAnalysisMessage struct {
	DetectionMessage[Detection]
}
