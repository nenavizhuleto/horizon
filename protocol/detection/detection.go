package detection

import "time"

type DetectionType string

// Primitive detection types
const (
	// Detect motion on video
	MotionDetection = DetectionType("motion")

	// Detect object on image/frame
	ObjectDetection = DetectionType("object")

	// Detect value on analog/digital sensor input
	ValueDetection = DetectionType("value")
)

type SensorType string

// Primitive sensor types
const (
	// IP-Camera / Videofile
	//
	// One may reason, how do video sensor differ from analog or digital?
	//
	// I don't know yet, but believe me.
	VideoSensor = SensorType("video")

	// Kind of sensor that used in object detection
	NeuralSensor = SensorType("neural")

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

type Detection struct {
	// Which sensor made detection
	Producer Sensor `json:"producer"`

	// Although each sensor has it's own type
	// and associated detection types with it,
	// it would be nice to pass detection type further,
	// for better API
	Type DetectionType `json:"detection_type"`

	// When did detection happened
	Timestamp time.Time `json:"timestamp"`

	// As detection value may vary based on different sensor types,
	// We should be able to pass any data to consumers.
	// So also we need some way to extract this value in sensor intended way.
	Value any `json:"value"`
}
