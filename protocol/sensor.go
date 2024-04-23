package protocol

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol/video"
)

type SensorType string

// Primitive sensor types
const (
	// IP-Camera / Videofile
	//
	// One may reason, how do video sensor differ from analog or digital?
	//
	// I don't know yet, but believe me.
	Motionsensor = SensorType("motion")

	// Kind of sensor that used in object detection
	ObjectSensor = SensorType("object")

	// Accelerometers
	// Light
	// Sound
	// Pressuer
	// Temperature
	// Humidity
	// Gas
	AnalogSensor  = SensorType("analog")
	DigitalSensor = SensorType("digital")
)

type Sensor struct {
	// Identifier
	ID string `json:"id"`

	// Reasonable sensor name
	Name string `json:"name"`

	// One of: digital | analog | video
	Type SensorType `json:"type"`
}

type ISensor interface {
	Sensor() Sensor
	SetID(id string)
	SetName(name string)
}

type IMotionSensor interface {
	ISensor
	NewMotion(src video.Source, pos video.Position) Motion
	NewMotionDetection(motions ...Motion) Detection
}

type IObjectSensor interface {
	ISensor
	NewObject(class string, bounding_box video.Position, conf float32) Object
	NewObjectDetection(objects ...Object) Detection
}

type IValueSensor interface {
	ISensor
	NewValueDetection(values ...Value) Detection
}

func (s *Sensor) SetID(id string) {
	s.ID = id
}

func (s *Sensor) SetName(name string) {
	s.Name = name
}

func (s Sensor) Sensor() Sensor {
	return s
}

func (s Sensor) NewMotion(src video.Source, pos video.Position) Motion {
	return Motion{
		Source:   src,
		Position: pos,
	}
}

func (s Sensor) NewMotionDetection(motions ...Motion) Detection {
	return Detection{
		Producer:  s,
		Type:      MotionDetection,
		Timestamp: time.Now(),
		Value:     motions,
	}
}

func (s Sensor) NewObject(class string, bounding_box video.Position, conf float32) Object {
	return Object{
		Class:       class,
		BoundingBox: bounding_box,
		Confidence:  conf,
	}
}

func (s Sensor) NewObjectDetection(objects ...Object) Detection {
	return Detection{
		Producer:  s,
		Type:      ObjectDetection,
		Timestamp: time.Now(),
		Value:     objects,
	}
}

func (s Sensor) NewValueDetection(values ...Value) Detection {
	return Detection{
		Producer:  s,
		Type:      ObjectDetection,
		Timestamp: time.Now(),
		Value:     values,
	}
}

func NewMotionSensor(id, name string) IMotionSensor {
	return Sensor{
		ID:   id,
		Name: name,
		Type: Motionsensor,
	}
}

func NewObjectSensor(id, name string) IObjectSensor {
	return Sensor{
		ID:   id,
		Name: name,
		Type: ObjectSensor,
	}
}
