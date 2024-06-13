package protocol

import "time"

type FrameLocation struct {
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Topic     string `json:"topic"`
}

type DetectionMessageBody[D Typable] struct {
	Timestamp     time.Time     `json:"timestamp"`
	FrameLocation FrameLocation `json:"frame_location"`
	Detections    []D           `json:"detections"`
}

func (m DetectionMessageBody[D]) Type() MessageType {
	var d D
	return Join(MessageDetection, d.Type())
}
