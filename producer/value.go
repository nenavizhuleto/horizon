package producer

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

type ValueDetectionProducer DetectionProducer[protocol.Value, protocol.ValueDetectionMessage, any]

type vdp protocol.ValueProducer

// NewValueDetectionProducer creates producer for value detection messages
func NewValueDetectionProducer(id, name string) ValueDetectionProducer {
	return vdp(protocol.NewValueProducer(id, name))
}

// NewDetectionMessage constructs new detection message
func (dp vdp) NewDetectionMessage(ts time.Time, options any, values ...protocol.Value) protocol.ValueDetectionMessage {
	return protocol.NewValueDetectionMessage(protocol.ValueProducer(dp), ts, values...)
}
