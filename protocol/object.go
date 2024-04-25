package protocol

import "time"

type Object struct {
	// Arbitrary value
	//
	// Example: person, car, ball, bycicle
	Class string `json:"class"`

	// Where object was detected
	BoundingBox Position `json:"bounding_box"`

	// How confident this detection is
	//
	// Floating-point number: 0.0 <= x <= 1.0
	Confidence float32 `json:"confidence"`

	// We also might want to have
	// additional information alongise with detection
	UserData any `json:"user_data"`
}

// Neural networks
type ObjectProducer struct {
	Producer
}

func NewObjectProducer(id, name string) ObjectProducer {
	return ObjectProducer{
		Producer: NewProducer(id, name),
	}
}

type FrameLocation struct {
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Topic     string `json:"topic"`
}

type ObjectDetectionOptions struct {
	FrameLocation *FrameLocation `json:"frame_location,omitempty"`
}

type ObjectDetectionMessage struct {
	ObjectDetectionOptions
	DetectionMessage[ObjectProducer]
	Objects []Object `json:"objects"`
}

func NewObjectDetectionMessage(producer ObjectProducer, ts time.Time, options ObjectDetectionOptions, objects ...Object) ObjectDetectionMessage {
	return ObjectDetectionMessage{
		DetectionMessage:       NewDetectionMessage(MessageObjectDetection, producer, ts),
		Objects:                objects,
		ObjectDetectionOptions: options,
	}
}

type ObjectAnalysisReport struct {
	ObjectDetectionOptions
	Report
	Object Object `json:"object"`
}

func NewObjectAnalysisMessage(producer AnalysisProducer, severity Severity, report ObjectAnalysisReport) AnalysisMessage[ObjectAnalysisReport] {
	return AnalysisMessage[ObjectAnalysisReport]{
		ProducerMessage: NewProducerMessage(MessageAnalysisObjectDetection, producer),
		Severity:        severity,
		Report:          report,
	}
}
