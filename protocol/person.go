package protocol

type PersonDetection struct {
	BoundingBox Position `json:"bounding_box"`
	Confidence  float32  `json:"confidence"`
}

type PersonDetectionMessage DetectionMessage[PersonDetection]

func NewPersonDetectionMessage(producer Producer, message PersonDetectionMessage) Message[PersonDetectionMessage] {
	return Message[PersonDetectionMessage]{
		Type:     MessagePersonDetection,
		Producer: producer,
		Body:     message,
	}
}

type PersonReport interface{}
type PersonAnalysis Analysis[PersonReport, PersonDetection]
type PersonAnalysisMessage AnalysisMessage[PersonAnalysis]

func NewPersonAnalysisMessage(producer Producer, message PersonAnalysisMessage) Message[PersonAnalysisMessage] {
	return Message[PersonAnalysisMessage]{
		Producer: producer,
		Type:     MessagePersonAnalysis,
		Body:     message,
	}
}
