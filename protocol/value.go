package protocol

import "time"

type Value Detection

// Accelerometers, Light, Sound, Pressuer, Temperature, Humidity, Gas
type ValueProducer struct {
	Producer
}

func NewValueProducer(id, name string) ValueProducer {
	return ValueProducer{
		Producer: NewProducer(id, name),
	}
}

type ValueDetectionMessage struct {
	DetectionMessage[ValueProducer]
	Values []Value `json:"values"`
}

func NewValueDetectionMessage(producer ValueProducer, ts time.Time, values ...Value) ValueDetectionMessage {
	return ValueDetectionMessage{
		DetectionMessage: NewDetectionMessage(MessageValueDetection, producer, ts),
		Values:           values,
	}
}

type ValueAnalysisReport struct {
	Report
	Value Value `json:"value"`
}

func NewValueAnalysisMessage(producer AnalysisProducer, severity Severity, report ValueAnalysisReport) AnalysisMessage[ValueAnalysisReport] {
	return AnalysisMessage[ValueAnalysisReport]{
		ProducerMessage: NewProducerMessage(MessageAnalysisValueDetection, producer),
		Severity:        severity,
		Report:          report,
	}
}
