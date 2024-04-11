package protocol

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
