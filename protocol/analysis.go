package protocol

type Severity string

const (
	SeverityInfo    = Severity("info")
	SeverityWarning = Severity("warning")
	SeverityPanic   = Severity("panic")
)

type AnalysisProducerOptions struct {
	MotionProducerOptions
	EventID *string `json:"event_id,omitempty"`
}

type AnalysisProducer struct {
	AnalysisProducerOptions
	Producer
}

func NewAnalysisProducer(id, name string, options AnalysisProducerOptions) AnalysisProducer {
	return AnalysisProducer{
		Producer:                NewProducer(id, name),
		AnalysisProducerOptions: options,
	}
}

type AnalysisMessage[T any] struct {
	ProducerMessage[AnalysisProducer]
	Severity Severity `json:"severity"`
	Report   T        `json:"report"`
}

func NewAnalysisMessage[T any](t MessageType, producer AnalysisProducer, severity Severity, report T) AnalysisMessage[T] {
	return AnalysisMessage[T]{
		ProducerMessage: NewProducerMessage(t, producer),
		Severity:        severity,
		Report:          report,
	}
}
