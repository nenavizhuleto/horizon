package protocol

type PlateDetection struct {
	BoundingBox Position `json:"bounding_box"`
	Confidence  float32  `json:"confidence"`
	Plate       string   `json:"plate"`
}

func (m PlateDetection) Type() MessageType {
	return "plate"
}

type PlateReport []List
type PlateDetectionMessageBody = DetectionMessageBody[PlateDetection]
type PlateAnalysesMessageBody = AnalysesMessageBody[PlateReport, PlateDetection]

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
