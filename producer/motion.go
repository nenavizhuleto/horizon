package producer

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

type MotionDetectionProducer DetectionProducer[protocol.Motion, protocol.MotionDetectionMessage, any]

type mdp protocol.MotionProducer

// NewMotionDetectionProducer creates producer for motion detection messages
func NewMotionDetectionProducer(id, name string, options protocol.MotionProducerOptions) MotionDetectionProducer {
	return mdp(protocol.NewMotionProducer(id, name, options))
}

// NewDetectionMessage constructs new detection message
func (dp mdp) NewDetectionMessage(ts time.Time, options any, motions ...protocol.Motion) protocol.MotionDetectionMessage {
	return protocol.NewMotionDetectionMessage(protocol.MotionProducer(dp), ts, motions...)
}
