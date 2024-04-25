package producer

import "github.com/nenavizhuleto/horizon/protocol"

// Used to produce valid frame messages
type FrameProducer interface {
	NewFrameMessage(frame protocol.Frame, options protocol.FrameMessageOptions) protocol.FrameMessage
}

type fp protocol.Producer

// NewFrameProducer creates producer for frame messages
func NewFrameProducer(id, name string) FrameProducer {
	return fp(protocol.NewProducer(id, name))
}

// NewFrameMessage constructs new frame message
func (p fp) NewFrameMessage(frame protocol.Frame, options protocol.FrameMessageOptions) protocol.FrameMessage {
	return protocol.NewFrameMessage(protocol.Producer(p), frame, options)
}
