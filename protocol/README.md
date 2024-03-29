# Protocol

## Input

- Detection

> What could be considered as `Detection`?

Any form of `algorithm`, `computer vision` or `neural network` output based on
input data collected from so called `sensors`

Example:

- Smoke level (algorithm)
- Motion (computer vision)
- Object detection (neural network)

```go
// Primitive detection types
const (
	// Detect motion on video
	MotionDetection = DetectionType("motion")

	// Detect object on image/frame
	ObjectDetection = DetectionType("object")

	// Detect value on analog/discrete sensor input
	ValueDetection = DetectionType("value")
)
```

> What does `sensor` mean?

Device that can produce specialized information about changes in surrounding environment

Example:

- IP-Camera
- Fire-detector
- Heat-detector
- Sound-detector
- Access Control System
- Other


```go
// Primitive sensor types
const (
	// IP-Camera / Videofile
	//
	// One may reason, how do video sensor differ from analog or digital?
	//
	// I don't know yet, but believe me.
	VideoSensor = SensorType("video")

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
```

```go
type Sensor struct {
	// Identifier
	ID string

	// Reasonable sensor name
	Name string

	// One of: digital | analog | video
	Type SensorType
}
```

### Detection specification

In order to distinguish one detection from another, we need some way to classify detections and mark them accourdingly

As said earlier, detections must be produced with sensors.

Each sensor should have it's own identifier and set of detection classes that it can produce

> What information every detection have to have?

We need to know:

- Who produced this detection?
- When detection is made?
- What kind of a detection this is?

```go
type Detection struct {
	// Which sensor made detection
	Producer Sensor

	// Although each sensor has it's own type
	// and associated detection types with it,
	// it would be nice to pass detection type further,
	// for better API
	Type DetectionType

	// When did detection happened
	Timestamp time.Time

	// As detection value may vary based on different sensor types,
	// We should be able to pass any data to consumers.
	// So also we need some way to extract this value in sensor intended way.
	Value any
}
```


#### Motion

Consider `Motion` detection

> Which `sensors` can produce `Motion` detection?

`Motion` corresponds directly to video processing which means 
that only sensor which is able to make this detection is some form of `Video`

By this definition, next sources can be considered as `Motion sensor`

- Any kind of video stream (RTMP, RTSP, etc.)
- Static video files (mp4, etc.)

Concrete implementation of this is on hands of sensor developers

> What information do we have to have in `Motion detection`

We should be able to answer to following questions:

- Who is the producer of this detection?
- When does this detection happend?
- What is the source upon which the detection was made on?
- (optional) Where exactly motion happend? (consider video frame boundary)

```go
type Motion struct {
	// Video stream or file
	Source VideoSource

	// Position where motion is detected (optional)
	Position *VideoPosition
}

type VideoSource struct {
	// Uniform Resource Identifier
	URI string

	// Width(X) and Height(Y) of the source, on which consumers could rely on
	Dimensions VideoDimensions
}

// Might used image.Point, but implemented for better marshal control
type VideoDimensions struct {
	Width  int
	Height int
}

// Might used image.Rectangle, but will not
type VideoPosition struct {
	X      int
	Y      int
	Width  int
	Height int
}
```

Example motion detection:

```go
Detection{
    Producer: Sensor{
        ID:   "xxxx-xxxx-xxxx-xxxx",
        Name: "my-beloved-camera",
        Type: VideoSensor,
    },
    Type:      MotionDetection,
    Timestamp: time.Now(),
    Value: Motion{
        Source: VideoSource{
            URI: "rtsp://localhost:8554/camera",
            Dimensions: VideoDimensions{
                Width:  1920,
                Height: 1080,
            },
        },
        Position: &VideoPosition{
            X:      200,
            Y:      400,
            Width:  500,
            Height: 300,
        },
    },
}
```

> Is `Producer` equal to `Sensor`?

`Producer` relates to concrete implementation of a `Sensor`.
Where `Sensor` is generalized interpretation of the thing

#### Object

To `Object` detection fall following kinds of image processing:

- Pattern recognition (license plates, pose)
- Object recognition (person, face, weapon, etc.)
- Image classification (fire, smoke)

This kind of detections are made by neural networks

> Should neural network be considered as sensor of some sort?

Actually, I would argue with "YES", but maybe this topic needs further investigations. So let's think of neural network as another sensor.

```go
// Kind of sensor that used in object detection
const NeuralSensor = SensorType("neural")
```

> What information do we need in `Object Detection`?

We need all information from base detection above and following ones:

- What class of object was detected? (arbitrary value aka. person, car, bicycle, etc.)
- Where it was detected? (bounding box inside image/frame)
- How confident sensor is about this detection? (floating point value between 0.0 and 1.0)

We also might want to have additional information alongside with detection

```go
type Object struct {
	// Arbitrary value
	//
	// Example: person, car, ball, bycicle
	Class string

	// Where object was detected
	BoundingBox video.Position

	// How confident this detection is
	//
	// Floating-point number: 0.0 <= x <= 1.0
	Confidence float32

	// We also might want to have
	// additional information alongise with detection
	UserData any
}
```

Example usage:

```go
Detection{
	Producer: Sensor{
		ID:   "xxxx-xxxx-xxxx-xxxx",
		Name: "license-plate-recognition",
		Type: NeuralSensor,
	},
	Type:      ObjectDetection,
	Timestamp: time.Now(),
	Value: Object{
		Class: "license plate",
		BoundingBox: video.Position{
			X:      300,
			Y:      200,
			Width:  300,
			Height: 300,
		},
		Confidence: 0.98,
		UserData: "AC_TU_AL_NU_MB_ER"
	},
}
```

## Output

> What analyses do we have to make upon detection?

1. We analyze `Event`
2. Upon certain `Detection`
3. And procude `Report`

```go
type Analysis struct {
	Event     Event
	Detection Detection
	Report    Report
}
```

```go
type Severity int

const (
	_ Severity = iota
	INFO
	WARN
	PANIC
)

type Report struct {
	UserData any
	Severity Severity
}
```

First of all, let's explain ourselves term: `Event`

an `Event` is continious period of time. It has start & end.
It has some `Trigger` from which `Event` is initially started.
Might have other data that we consumed in-event-time

```go
type Trigger Detection

type Event struct {
	ID      string
	Trigger Trigger

	Start    time.Time
	End      time.Time
	Duration time.Duration
}
```



During the `Event` we want to analyze detections as they come.

What we can analyze?

If it's a `Motion` Detection:

- Do we allow for this specific `Sensor` to produce this `Detection` in current time?

> sidenote: can we somehow use here linked-list?

> When new event is started?

We need some sort of `Trigger` to fire an `Event`. 
By domain, we assume that some `Detection` might be the `Trigger.

Let's consider `Motion` detection as `Trigger` to start an `Event`.
So, when we consume `Motion` detection, following process can arise:

1. Record this detection as `Trigger`
2. Record event start timetstamp and start timer





- Analysis
