package producer

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

type MotionDetectionProducer DetectionProducer[protocol.Motion, protocol.MotionDetectionMessage]

type mdp protocol.MotionProducer

// NewMotionDetectionProducer creates producer for motion detection messages
func NewMotionDetectionProducer(id, name string) MotionDetectionProducer {
	mp := mdp(protocol.NewMotionProducer(id, name))
	return mp
}

// NewDetectionMessage constructs new detection message
func (dp mdp) NewDetectionMessage(ts time.Time, motions ...protocol.Motion) protocol.MotionDetectionMessage {
	return protocol.NewMotionDetectionMessage(protocol.MotionProducer(dp), ts, motions...)
}
