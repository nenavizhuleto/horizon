package producer

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

type EventProducer interface {
	NewStartMessage(id string, start time.Time, trigger protocol.Detection) protocol.EventStartMessage
	NewEndMessage(id string, end time.Time, duration time.Duration) protocol.EventEndMessage
}

type ep protocol.EventProducer

func NewEventProducer(id, name string, options protocol.EventProducerOptions) EventProducer {
	return ep(protocol.NewEventProducer(id, name, options))
}

func (p ep) NewStartMessage(id string, start time.Time, trigger protocol.Detection) protocol.EventStartMessage {
	return protocol.NewEventStartMessage(protocol.EventProducer(p), id, start, trigger)
}

func (p ep) NewEndMessage(id string, end time.Time, duration time.Duration) protocol.EventEndMessage {
	return protocol.NewEventEndMessage(protocol.EventProducer(p), id, end, duration)
}
