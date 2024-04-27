package producer

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

type MessageProducer interface {
	NewObjectDetectionMessage(ts time.Time, objects []protocol.Object, options ...protocol.MessageOptions) protocol.ObjectDetectionMessage
	NewMotionDetectionMessage(ts time.Time, motions []protocol.Motion, options ...protocol.MessageOptions) protocol.MotionDetectionMessage
	NewValueDetectionMessage(ts time.Time, values []protocol.Value, options ...protocol.MessageOptions) protocol.ValueDetectionMessage

	NewFrameMessage(frame protocol.Frame, options ...protocol.MessageOptions) protocol.FrameMessage

	NewObjectAnalysisMessage(ts time.Time, severity protocol.Severity, analyses []protocol.ObjectAnalysis, options ...protocol.MessageOptions) protocol.AnalysisMessage[protocol.Object]
	NewMotionAnalysisMessage(ts time.Time, severity protocol.Severity, analyses []protocol.MotionAnalysis, options ...protocol.MessageOptions) protocol.AnalysisMessage[protocol.Motion]
	NewValueAnalysisMessage(ts time.Time, severity protocol.Severity, analyses []protocol.ValueAnalysis, options ...protocol.MessageOptions) protocol.AnalysisMessage[protocol.Value]

	NewEventStartMessage(id string, start time.Time, trigger protocol.Detection, options ...protocol.MessageOptions) protocol.EventStartMessage
	NewEventEndMessage(id string, end time.Time, duration time.Duration, options ...protocol.MessageOptions) protocol.EventEndMessage
}

type mp protocol.Producer

func NewMessageProducer(id, name string, options ...protocol.ProducerOptions) MessageProducer {
	return mp(protocol.NewProducer(id, name, options...))
}

func (p mp) NewObjectDetectionMessage(ts time.Time, objects []protocol.Object, options ...protocol.MessageOptions) protocol.ObjectDetectionMessage {
	return protocol.NewObjectDetectionMessage(protocol.Producer(p), ts, objects, options...)
}
func (p mp) NewMotionDetectionMessage(ts time.Time, motions []protocol.Motion, options ...protocol.MessageOptions) protocol.MotionDetectionMessage {
	return protocol.NewMotionDetectionMessage(protocol.Producer(p), ts, motions, options...)
}
func (p mp) NewValueDetectionMessage(ts time.Time, values []protocol.Value, options ...protocol.MessageOptions) protocol.ValueDetectionMessage {
	return protocol.NewValueDetectionMessage(protocol.Producer(p), ts, values, options...)
}
func (p mp) NewFrameMessage(frame protocol.Frame, options ...protocol.MessageOptions) protocol.FrameMessage {
	return protocol.NewFrameMessage(protocol.Producer(p), frame, options...)
}

func (p mp) NewObjectAnalysisMessage(ts time.Time, severity protocol.Severity, analyses []protocol.ObjectAnalysis, options ...protocol.MessageOptions) protocol.AnalysisMessage[protocol.Object] {
	return protocol.NewAnalysisMessage(protocol.Producer(p), ts, severity, analyses, options...)
}
func (p mp) NewMotionAnalysisMessage(ts time.Time, severity protocol.Severity, analyses []protocol.MotionAnalysis, options ...protocol.MessageOptions) protocol.AnalysisMessage[protocol.Motion] {
	return protocol.NewAnalysisMessage(protocol.Producer(p), ts, severity, analyses, options...)
}
func (p mp) NewValueAnalysisMessage(ts time.Time, severity protocol.Severity, analyses []protocol.ValueAnalysis, options ...protocol.MessageOptions) protocol.AnalysisMessage[protocol.Value] {
	return protocol.NewAnalysisMessage(protocol.Producer(p), ts, severity, analyses, options...)
}

func (p mp) NewEventStartMessage(id string, start time.Time, trigger protocol.Detection, options ...protocol.MessageOptions) protocol.EventStartMessage {
	return protocol.NewEventStartMessage(protocol.Producer(p), id, start, trigger, options...)
}
func (p mp) NewEventEndMessage(id string, end time.Time, duration time.Duration, options ...protocol.MessageOptions) protocol.EventEndMessage {
	return protocol.NewEventEndMessage(protocol.Producer(p), id, end, duration, options...)
}
