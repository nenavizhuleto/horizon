package protocol

type ObjectDetection struct {
	ID          string   `json:"id"`
	Class       string   `json:"class"`
	BoundingBox Position `json:"bounding_box"`
	Confidence  float32  `json:"confidence"`
}

type ObjectDetectionMessage = DetectionMessage[ObjectDetection]

func (p ObjectDetection) Type() MessageType {
	return MessageObject
}

func NewObjectDetectionMessage(producer Camera, message ObjectDetectionMessage) Message[ObjectDetectionMessage] {
	return Message[ObjectDetectionMessage]{
		Type:     MessageObjectDetection,
		Producer: producer,
		Body:     message,
	}
}

type ObjectReport interface{}
type ObjectAnalysis Analysis[ObjectReport, ObjectDetection]
type ObjectAnalysesMessage AnalysesMessage[ObjectAnalysis]

func NewObjectAnalysesMessage(producer Camera, message ObjectAnalysesMessage) Message[ObjectAnalysesMessage] {
	return Message[ObjectAnalysesMessage]{
		Type:     MessageObjectAnalysis,
		Producer: producer,
		Body:     message,
	}
}
