package producer

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

// Used to produce valid extended object detection messages
type ObjectDetectionProducer DetectionProducer[protocol.Object, protocol.ObjectDetectionMessage, protocol.ObjectDetectionOptions]

type odp protocol.ObjectProducer

// NewObjectDetectionProducer creates producer for object detection messages with extended capabilities
func NewObjectDetectionProducer(id, name string) ObjectDetectionProducer {
	op := odp(protocol.NewObjectProducer(id, name))
	return op
}

// NewExtDetectionMessage constructs new detection message with extended options
func (dp odp) NewDetectionMessage(ts time.Time, options protocol.ObjectDetectionOptions, objects ...protocol.Object) protocol.ObjectDetectionMessage {
	return protocol.NewObjectDetectionMessage(protocol.ObjectProducer(dp), ts, options, objects...)
}
