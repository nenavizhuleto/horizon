package protocol

import "time"

type Detection interface{}

type DetectionMessage[T any] struct {
	ProducerMessage[T]
	Timestamp  time.Time   `json:"timestamp"`
	Detections []Detection `json:"detections"`
}

func NewDetectionMessage[T any](t MessageType, producer T, ts time.Time) DetectionMessage[T] {
	return DetectionMessage[T]{
		ProducerMessage: NewProducerMessage(t, producer),
		Timestamp:       ts,
	}
}

type Motion struct {
	// Video stream or file
	Source Source `json:"source"`

	// Position where motion is detected
	Position Position `json:"position,omitempty"`
}

type MotionDetectionMessage struct {
	DetectionMessage[MotionProducer]
	Detections []Motion `json:"detections"`
}

func NewMotionDetectionMessage(producer MotionProducer, ts time.Time, motions ...Motion) MotionDetectionMessage {
	return MotionDetectionMessage{
		DetectionMessage: NewDetectionMessage(MessageMotionDetection, producer, ts),
		Detections:       motions,
	}
}

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

type ObjectDetectionMessage struct {
	DetectionMessage[ObjectProducer]
	Detections []Object `json:"detections"`
}

type FrameLocation struct {
	Partition int32  `json:"partition"`
	Offset    int64  `json:"offset"`
	Topic     string `json:"topic"`
}

type ObjectDetectionExtOptions struct {
	FrameLocation *FrameLocation `json:"frame_location,omitempty"`
}

type ObjectDetectionExtMessage struct {
	ObjectDetectionExtOptions
	DetectionMessage[ObjectProducer]
	Detections []Object `json:"detections"`
}

func NewObjectDetectionMessage(producer ObjectProducer, ts time.Time, objects ...Object) ObjectDetectionMessage {
	return ObjectDetectionMessage{
		DetectionMessage: NewDetectionMessage(MessageObjectDetection, producer, ts),
		Detections:       objects,
	}
}

func NewObjectDetectionExtMessage(producer ObjectProducer, ts time.Time, options ObjectDetectionExtOptions, objects ...Object) ObjectDetectionExtMessage {
	return ObjectDetectionExtMessage{
		DetectionMessage:          NewDetectionMessage(MessageObjectDetectionExt, producer, ts),
		Detections:                objects,
		ObjectDetectionExtOptions: options,
	}
}

type Value any

type ValueDetectionMessage struct {
	DetectionMessage[ValueProducer]
	Detections []Value `json:"detections"`
}

func NewValueDetectionMessage(producer ValueProducer, ts time.Time, values ...Value) ValueDetectionMessage {
	return ValueDetectionMessage{
		DetectionMessage: NewDetectionMessage(MessageValueDetection, producer, ts),
		Detections:       values,
	}
}
