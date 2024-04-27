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

type EventProducerOptions struct {
	Camera *Camera `json:"camera,omitempty"`
}

type EventProducer struct {
	EventProducerOptions
	Producer
}

type EventMessage struct {
	ProducerMessage[EventProducer]
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

func NewEventProducer(id, name string, options EventProducerOptions) EventProducer {
	return EventProducer{
		Producer:             NewProducer(id, name),
		EventProducerOptions: options,
	}
}

func NewEventMessage(t MessageType, producer EventProducer, id string) EventMessage {
	return EventMessage{
		ProducerMessage: NewProducerMessage(t, producer),
		ID:              id,
	}
}

func NewEventStartMessage(producer EventProducer, id string, start time.Time, trigger Detection) EventStartMessage {
	return EventStartMessage{
		EventMessage: NewEventMessage(MessageEventStart, producer, id),
		Start:        start,
		Trigger:      trigger,
	}
}

func NewEventEndMessage(producer EventProducer, id string, end time.Time, duration time.Duration) EventEndMessage {
	return EventEndMessage{
		EventMessage: NewEventMessage(MessageEventEnd, producer, id),
		End:          end,
		Duration:     duration,
	}
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
