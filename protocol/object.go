package protocol

import "time"

type FrameLocation struct {
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Topic     string `json:"topic"`
}

type DetectionMessage[D any] struct {
	Timestamp     time.Time     `json:"timestamp"`
	FrameLocation FrameLocation `json:"frame_location"`
	Detections    []D           `json:"detections"`
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
type ObjectAnalysesMessage AnalysesMessage[ObjectAnalysis]

func NewObjectAnalysesMessage(producer Producer, message ObjectAnalysesMessage) Message[ObjectAnalysesMessage] {
	return Message[ObjectAnalysesMessage]{
		Type:     MessageObjectAnalysis,
		Producer: producer,
		Body:     message,
	}
}
