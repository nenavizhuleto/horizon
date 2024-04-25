package producer_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/nenavizhuleto/horizon/producer"
	"github.com/nenavizhuleto/horizon/protocol"
)

var (
	id   = "id"
	name = "producer"
)

func PrintMessage(t *testing.T, msg any) {
	bytes, err := json.MarshalIndent(msg, "", "  ")

	if err != nil {
		panic(err)
	}

	t.Logf(string(bytes))
}

func Test_MotionDetectionProducer(t *testing.T) {

	mdp := producer.NewMotionDetectionProducer(id, name, protocol.MotionProducerOptions{
		Camera: &protocol.Camera{
			Name:  "fuck",
			Group: "you",
		},
	})

	msg := mdp.NewDetectionMessage(time.Now(), nil, protocol.Motion{}, protocol.Motion{})

	PrintMessage(t, msg)
}

func Test_ObjectDetectionProducer(t *testing.T) {
	odp := producer.NewObjectDetectionProducer(id, name)

	msg := odp.NewDetectionMessage(time.Now(), protocol.ObjectDetectionOptions{
		FrameLocation: &protocol.FrameLocation{
			Partition: 123,
			Offset:    332,
		},
	}, protocol.Object{Class: "person"}, protocol.Object{Class: "weapon"}, protocol.Object{Class: "plate"})

	PrintMessage(t, msg)
}

func Test_ValueDetectionProducer(t *testing.T) {
	vdp := producer.NewValueDetectionProducer(id, name)

	msg := vdp.NewDetectionMessage(time.Now(), nil, 123, 555, "hello", "omh", true, struct{}{})

	PrintMessage(t, msg)
}

func Test_FrameProducer(t *testing.T) {
	fp := producer.NewFrameProducer(id, name)

	regs := []protocol.Position{
		{X: 0, Y: 0, Width: 100, Height: 200},
	}

	msg := fp.NewFrameMessage([]byte{0, 1, 2, 4, 5, 6, 7, 8}, protocol.FrameMessageOptions{
		Dimensions: &protocol.Dimensions{
			Width:  100,
			Height: 2333,
		},
		Regions: &regs,
	})

	PrintMessage(t, msg)
}

func Test_MotionAnalysisProducer(t *testing.T) {
	manp := producer.NewMotionAnalysisProducer(id, name, protocol.AnalysisProducerOptions{
		MotionProducerOptions: protocol.MotionProducerOptions{
			Camera: &protocol.Camera{
				Name:  "fuck",
				Group: "you",
			},
		},
	})

	msg := manp.NewAnalysisMessage(protocol.SeverityWarning, protocol.MotionAnalysisReport{
		Report: protocol.Report{
			Producer:  protocol.NewProducer("new_id", "new_producer"),
			Timestamp: time.Now(),
		},
		Motion: protocol.Motion{
			Source: protocol.Source{
				URI: "helloworld",
			},
		},
	})

	PrintMessage(t, msg)
}

func Test_ObjectAnalysisProducer(t *testing.T) {
	oanp := producer.NewObjectAnalysisProducer(id, name, protocol.AnalysisProducerOptions{})

	msg := oanp.NewAnalysisMessage(protocol.SeverityPanic, protocol.ObjectAnalysisReport{
		ObjectDetectionOptions: protocol.ObjectDetectionOptions{
			FrameLocation: &protocol.FrameLocation{
				Partition: 1234,
				Offset:    1234,
				Topic:     "topic",
			},
		},
		Report: protocol.Report{
			Producer:  protocol.NewProducer("", ""),
			Timestamp: time.Now(),
			UserData: map[string]any{
				"name":      "john",
				"direction": "forward",
			},
		},
		Object: protocol.Object{
			Class: "person",
		},
	})

	PrintMessage(t, msg)
}

func Test_ValueAnalysisProducer(t *testing.T) {
	vanp := producer.NewValueAnalysisProducer(id, name, protocol.AnalysisProducerOptions{})

	msg := vanp.NewAnalysisMessage(protocol.SeverityInfo, protocol.ValueAnalysisReport{
		Report: protocol.Report{
			Producer:  protocol.NewProducer("", ""),
			Timestamp: time.Now(),
		},
		Value: 1259812,
	})

	PrintMessage(t, msg)
}

func Test_EventAnalysisProducer(t *testing.T) {
	eanp := producer.NewEventAnalysisProducer(id, name, protocol.AnalysisProducerOptions{})

	msg := eanp.NewAnalysisMessage(protocol.SeverityPanic, protocol.EventAnalysisReport{
		Report: protocol.Report{
			Producer:  protocol.NewProducer("", ""),
			Timestamp: time.Now(),
		},
		Event: protocol.Event{
			ID: "event_id",
		},
		Detections: []protocol.Detection{
			protocol.Motion{},
			protocol.Object{},
			123,
			protocol.Motion{},
		},
	})

	PrintMessage(t, msg)
}
