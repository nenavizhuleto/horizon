package protocol

import (
	"time"
)

type Event struct {
	ID       string        `json:"id"`
	Trigger  Detection     `json:"trigger"`
	Start    time.Time     `json:"start"`
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
}

type EventMessage struct {
	ProducerMessage[Producer]
	ID string `json:"id"`
}

type EventStartMessage struct {
	EventMessage
	Trigger Detection `json:"trigger"`
	Start   time.Time `json:"start"`
}

type EventEndMessage struct {
	EventMessage
	End      time.Time     `json:"end"`
	Duration time.Duration `json:"duration"`
}

type EventAnalysisReport struct {
	Report
	Event      Event       `json:"event"`
	Detections []Detection `json:"detections"`
}

func NewEventAnalysisMessage(producer AnalysisProducer, severity Severity, report EventAnalysisReport) AnalysisMessage[EventAnalysisReport] {
	return AnalysisMessage[EventAnalysisReport]{
		ProducerMessage: NewProducerMessage(MessageAnalysisEvent, producer),
		Severity:        severity,
		Report:          report,
	}
}
