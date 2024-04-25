package protocol

import "time"

type Motion struct {
	// Video stream or file
	Source Source `json:"source"`

	// Position where motion is detected
	Position Position `json:"position,omitempty"`
}

type Camera struct {
	Name  string `json:"name"`
	Group string `json:"group"`
}

type MotionProducerOptions struct {
	Camera *Camera `json:"camera,omitempty"`
}

type MotionProducer struct {
	MotionProducerOptions
	Producer
}

func NewMotionProducer(id, name string, options MotionProducerOptions) MotionProducer {
	return MotionProducer{
		Producer:              NewProducer(id, name),
		MotionProducerOptions: options,
	}
}

type MotionDetectionMessage struct {
	DetectionMessage[MotionProducer]
	Motions []Motion `json:"motions"`
}

func NewMotionDetectionMessage(producer MotionProducer, ts time.Time, motions ...Motion) MotionDetectionMessage {
	return MotionDetectionMessage{
		DetectionMessage: NewDetectionMessage(MessageMotionDetection, producer, ts),
		Motions:          motions,
	}
}

type MotionAnalysisReport struct {
	Report
	Motion Motion `json:"motion"`
}

func NewMotionAnalysisMessage(producer AnalysisProducer, severity Severity, report MotionAnalysisReport) AnalysisMessage[MotionAnalysisReport] {
	return AnalysisMessage[MotionAnalysisReport]{
		ProducerMessage: NewProducerMessage(MessageAnalysisMotionDetection, producer),
		Severity:        severity,
		Report:          report,
	}
}
