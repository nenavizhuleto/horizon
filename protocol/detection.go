package protocol

import "time"

type Detection interface {
	PlateDetection | ObjectDetection
	Type() MessageType
}

type FrameLocation struct {
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Topic     string `json:"topic"`
}

type DetectionMessage[D Detection] struct {
	Timestamp     time.Time     `json:"timestamp"`
	FrameLocation FrameLocation `json:"frame_location"`
	Detections    []D           `json:"detections"`
}

func (m DetectionMessage[D]) Type() MessageType {
	var d D
	return NewMessageType(MessageDetection, d.Type())
}

func NewDetectionMessage[D Detection](detections []D) DetectionMessage[D] {
	return DetectionMessage[D]{
		Detections: detections,
	}
}
