package protocol_test

import (
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
	"github.com/nenavizhuleto/horizon/protocol/video"
)

func Test_MotionSensor(t *testing.T) {
	sensor := protocol.NewMotionSensor("id", "motion")

	motion := sensor.NewMotion(video.Source{}, video.Position{})

	actual := sensor.NewMotionDetection(motion)

	ts := time.Now()

	actual.SetTimestamp(ts)

	target := protocol.Detection{
		Producer:  sensor.Sensor(),
		Timestamp: ts,
		Type:      protocol.MotionDetection,
		Value: []protocol.Motion{
			{
				Source:   video.Source{},
				Position: video.Position{},
			},
		},
	}

	if !reflect.DeepEqual(target, actual) {
		log.Fatalf("target != actual\ntarget: %+v\nactual: %+v\n", target, actual)
	}

}

func Test_ObjectSensor(t *testing.T) {
	sensor := protocol.NewObjectSensor("id", "motion")

	object := sensor.NewObject("person", video.Position{}, 0.98)

	object.SetUserData("john")

	actual := sensor.NewObjectDetection(object)

	ts := time.Now()

	actual.SetTimestamp(ts)

	target := protocol.Detection{
		Producer:  sensor.Sensor(),
		Timestamp: ts,
		Type:      protocol.ObjectDetection,
		Value: []protocol.Object{
			{
				Class:       "person",
				BoundingBox: video.Position{},
				Confidence:  0.98,
				UserData:    "john",
			},
		},
	}

	if !reflect.DeepEqual(target, actual) {
		log.Fatalf("target != actual\ntarget: %+v\nactual: %+v\n", target, actual)
	}

}
