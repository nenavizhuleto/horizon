package protocol

type ObjectDetection struct {
	ID          string   `json:"id"`
	Class       string   `json:"class"`
	BoundingBox Position `json:"bounding_box"`
	Confidence  float32  `json:"confidence"`
}

func (d ObjectDetection) Type() MessageType {
	return "object"
}

type ObjectReport interface{}
type ObjectDetectionMessageBody = DetectionMessageBody[ObjectDetection]
type ObjectAnalysesMessageBody = AnalysesMessageBody[ObjectReport, ObjectDetection]
