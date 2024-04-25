package producer

import "github.com/nenavizhuleto/horizon/protocol"

type AnalysisProducer[R any] interface {
	NewAnalysisMessage(severity protocol.Severity, report R) protocol.AnalysisMessage[R]
}

type MotionAnalysisProducer AnalysisProducer[protocol.MotionAnalysisReport]

// MotionAnalysisProducer
type manp protocol.AnalysisProducer

func (p manp) NewAnalysisMessage(severity protocol.Severity, report protocol.MotionAnalysisReport) protocol.AnalysisMessage[protocol.MotionAnalysisReport] {
	return protocol.NewMotionAnalysisMessage(protocol.AnalysisProducer(p), severity, report)
}

func NewMotionAnalysisProducer(id, name string, options protocol.AnalysisProducerOptions) MotionAnalysisProducer {
	return manp(protocol.NewAnalysisProducer(id, name, options))
}

type ObjectAnalysisProducer AnalysisProducer[protocol.ObjectAnalysisReport]

// ObjectAnalysisProducer
type oanp protocol.AnalysisProducer

func (p oanp) NewAnalysisMessage(severity protocol.Severity, report protocol.ObjectAnalysisReport) protocol.AnalysisMessage[protocol.ObjectAnalysisReport] {
	return protocol.NewObjectAnalysisMessage(protocol.AnalysisProducer(p), severity, report)
}

func NewObjectAnalysisProducer(id, name string, options protocol.AnalysisProducerOptions) ObjectAnalysisProducer {
	return oanp(protocol.NewAnalysisProducer(id, name, options))
}

type ValueAnalysisProducer AnalysisProducer[protocol.ValueAnalysisReport]

// ValueAnalysisProducer
type vanp protocol.AnalysisProducer

func (p vanp) NewAnalysisMessage(severity protocol.Severity, report protocol.ValueAnalysisReport) protocol.AnalysisMessage[protocol.ValueAnalysisReport] {
	return protocol.NewValueAnalysisMessage(protocol.AnalysisProducer(p), severity, report)
}

func NewValueAnalysisProducer(id, name string, options protocol.AnalysisProducerOptions) ValueAnalysisProducer {
	return vanp(protocol.NewAnalysisProducer(id, name, options))
}

type EventAnalysisProducer AnalysisProducer[protocol.EventAnalysisReport]

// EventAnalysisProducer
type eanp protocol.AnalysisProducer

func (p eanp) NewAnalysisMessage(severity protocol.Severity, report protocol.EventAnalysisReport) protocol.AnalysisMessage[protocol.EventAnalysisReport] {
	return protocol.NewEventAnalysisMessage(protocol.AnalysisProducer(p), severity, report)
}

func NewEventAnalysisProducer(id, name string, options protocol.AnalysisProducerOptions) EventAnalysisProducer {
	return eanp(protocol.NewAnalysisProducer(id, name, options))
}
