package producer

import "github.com/nenavizhuleto/horizon/protocol"

// Used to produce valid frame messages
type FrameProducer interface {
	NewFrameMessage(frame protocol.Frame) protocol.FrameMessage
	NewExtFrameMessage(frame protocol.Frame, dims protocol.Dimensions, regs ...protocol.Position) protocol.ExtFrameMessage
}

type fp protocol.Producer

// NewFrameProducer creates producer for frame messages
func NewFrameProducer(id, name string) FrameProducer {
	return fp(protocol.NewProducer(id, name))
}

// NewFrameMessage constructs new frame message
func (p fp) NewFrameMessage(frame protocol.Frame) protocol.FrameMessage {
	return protocol.NewFrameMessage(protocol.Producer(p), frame)
}

// NewFrameExtMessage constructs new frame message
func (p fp) NewExtFrameMessage(frame protocol.Frame, dims protocol.Dimensions, regs ...protocol.Position) protocol.ExtFrameMessage {
	return protocol.NewExtFrameMessage(protocol.Producer(p), frame, dims, regs...)
}
