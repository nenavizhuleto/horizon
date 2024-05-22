package protocol

type PlateDetection struct {
	BoundingBox Position `json:"bounding_box"`
	Confidence  float32  `json:"confidence"`
	Plate       string   `json:"plate"`
}

type PlateDetectionMessage DetectionMessage[PlateDetection]

func NewPlateDetectionMessage(producer Producer, message PlateDetectionMessage) Message[PlateDetectionMessage] {
	return Message[PlateDetectionMessage]{
		Producer: producer,
		Type:     MessagePlateDetection,
		Body:     message,
	}
}

type Vehicle struct {
	ID    string `json:"id"`
	Plate string `json:"plate"`
	Brand string `json:"brand,omitempty"`
	Model string `json:"model,omitempty"`
	Color string `json:"color,omitempty"`
}

type List struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Severity Severity `json:"severity"`
	Color    string   `json:"color"`
}

type PlateReport []List
type PlateAnalysis Analysis[PlateReport, PlateDetection]
type PlateAnalysesMessage AnalysesMessage[PlateAnalysis]

func NewPlateAnalysesMessage(producer Producer, message PlateAnalysesMessage) Message[PlateAnalysesMessage] {
	return Message[PlateAnalysesMessage]{
		Producer: producer,
		Type:     MessagePlateAnalysis,
		Body:     message,
	}
}
