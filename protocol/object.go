package protocol

import "time"

type Location struct {
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Topic     string `json:"topic"`
}

type DetectionMessage[D any] struct {
	Timestamp  time.Time `json:"timestamp"`
	Location   Location  `json:"location"`
	Detections []D       `json:"detections"`
}

type ObjectDetection struct {
	ID          string   `json:"id"`
	Class       string   `json:"class"`
	BoundingBox Position `json:"bounding_box"`
	Confidence  float32  `json:"confidence"`
}

type ObjectDetectionMessage DetectionMessage[ObjectDetection]

func NewObjectDetectionMessage(producer Producer, message ObjectDetectionMessage) Message[ObjectDetectionMessage] {
	return Message[ObjectDetectionMessage]{
		Type:     MessageObjectDetection,
		Producer: producer,
		Body:     message,
	}
}

type ObjectReport interface{}
type ObjectAnalysis Analysis[ObjectReport, ObjectDetection]
type ObjectAnalysisMessage AnalysisMessage[ObjectAnalysis]

func NewObjectAnalysisMessage(producer Producer, message ObjectAnalysisMessage) Message[ObjectAnalysisMessage] {
	return Message[ObjectAnalysisMessage]{
		Type:     MessageObjectAnalysis,
		Producer: producer,
		Body:     message,
	}
}
