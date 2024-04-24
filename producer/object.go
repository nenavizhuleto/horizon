package producer

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

// Used to produce valid object detection messages
type ObjectDetectionProducer DetectionProducer[protocol.Object, protocol.ObjectDetectionMessage]

// Used to produce valid extended object detection messages
type ObjectDetectionExtProducer DetectionExtProducer[protocol.Object, protocol.ObjectDetectionExtMessage, protocol.ObjectDetectionExtOptions]

type odp protocol.ObjectProducer

// NewObjectDetectionProducer creates producer for object detection messages
func NewObjectDetectionProducer(id, name string) ObjectDetectionProducer {
	op := odp(protocol.NewObjectProducer(id, name))
	return op
}

// NewObjectDetectionExtProducer creates producer for object detection messages with extended capabilities
func NewObjectDetectionExtProducer(id, name string) ObjectDetectionExtProducer {
	op := odp(protocol.NewObjectProducer(id, name))
	return &op
}

// NewDetectionMessage constructs new detection message
func (dp odp) NewDetectionMessage(ts time.Time, objects ...protocol.Object) protocol.ObjectDetectionMessage {
	return protocol.NewObjectDetectionMessage(protocol.ObjectProducer(dp), ts, objects...)
}

// NewExtDetectionMessage constructs new detection message with extended options
func (dp odp) NewExtDetectionMessage(ts time.Time, options protocol.ObjectDetectionExtOptions, objects ...protocol.Object) protocol.ObjectDetectionExtMessage {
	return protocol.NewObjectDetectionExtMessage(protocol.ObjectProducer(dp), ts, options, objects...)
}
