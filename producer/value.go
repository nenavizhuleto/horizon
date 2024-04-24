package producer

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

type ValueDetectionProducer DetectionProducer[protocol.Value, protocol.ValueDetectionMessage]

type vdp protocol.ValueProducer

// NewValueDetectionProducer creates producer for value detection messages
func NewValueDetectionProducer(id, name string) ValueDetectionProducer {
	vp := vdp(protocol.NewValueProducer(id, name))
	return vp
}

// NewDetectionMessage constructs new detection message
func (dp vdp) NewDetectionMessage(ts time.Time, values ...protocol.Value) protocol.ValueDetectionMessage {
	return protocol.NewValueDetectionMessage(protocol.ValueProducer(dp), ts, values...)
}
