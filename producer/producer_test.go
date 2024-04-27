package producer_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/nenavizhuleto/horizon/producer"
	"github.com/nenavizhuleto/horizon/protocol"
)

var (
	ts = time.Now()
	p  = producer.NewMessageProducer("PRODUCER_ID", "PRODUCER_NAME", protocol.ProducerOptions{
		Camera: &protocol.Camera{
			Name:  "CAMERA_NAME",
			Group: "CAMERA_GROUP",
		},
	})

	event = protocol.Event{
		ID:       "event-id",
		Trigger:  protocol.Motion{},
		Start:    ts,
		End:      ts,
		Duration: 10 * time.Second,
	}
)

func PrintMessage(t *testing.T, msg any) {
	bytes, err := json.MarshalIndent(msg, "", "  ")

	if err != nil {
		panic(err)
	}

	t.Logf(string(bytes))
}

func Test_MotionDetectionProducer(t *testing.T) {
	msg := p.NewMotionDetectionMessage(ts, []protocol.Motion{{}, {}}, protocol.MessageOptions{
		Frame: &protocol.FrameOptions{
			Regions: &[]protocol.Position{
				{},
				{},
			},
			Dimensions: &protocol.Dimensions{
				Width:  1000,
				Height: 2000,
			},
		},
	})

	PrintMessage(t, msg)
}

func Test_ObjectDetectionProducer(t *testing.T) {

	msg := p.NewObjectDetectionMessage(ts, []protocol.Object{
		protocol.Object{Class: "person"}, protocol.Object{Class: "weapon"}, protocol.Object{Class: "plate"},
	}, protocol.MessageOptions{
		Frame: &protocol.FrameOptions{
			Location: &protocol.Location{
				Partition: 123,
				Offset:    332,
				Topic:     "TOPIC_NAME",
			},
		},
	})

	PrintMessage(t, msg)
}

func Test_ValueDetectionProducer(t *testing.T) {
	msg := p.NewValueDetectionMessage(ts, []protocol.Value{123, 555, "hello", "omh", true, struct{}{}})

	PrintMessage(t, msg)
}

func Test_FrameProducer(t *testing.T) {
	regs := []protocol.Position{
		{X: 0, Y: 0, Width: 100, Height: 200},
	}

	msg := p.NewFrameMessage([]byte{0, 1, 2, 4, 5, 6, 7, 8}, protocol.MessageOptions{
		Frame: &protocol.FrameOptions{
			Dimensions: &protocol.Dimensions{
				Width:  100,
				Height: 2333,
			},
			Regions: &regs,
		},
	})

	PrintMessage(t, msg)
}

func Test_EventProducer(t *testing.T) {

	start_msg := p.NewEventStartMessage(event.ID, event.Start, event.Trigger)
	end_msg := p.NewEventEndMessage(event.ID, event.End, event.Duration)

	PrintMessage(t, start_msg)

	PrintMessage(t, end_msg)

}

func Test_MotionAnalysisProducer(t *testing.T) {
	msg := p.NewMotionAnalysisMessage(ts, protocol.SeverityWarning, []protocol.MotionAnalysis{
		protocol.NewAnalysis(protocol.Motion{}, "this motion is too fast"),
	}, protocol.MessageOptions{
		Event: &protocol.EventOptions{
			ID: &event.ID,
		},
	})

	PrintMessage(t, msg)
}

func Test_ObjectAnalysisProducer(t *testing.T) {
	msg := p.NewObjectAnalysisMessage(ts, protocol.SeverityPanic, []protocol.ObjectAnalysis{
		protocol.NewAnalysis(protocol.Object{Class: "person"}, map[string]any{
			"name":      "john",
			"direction": "forward",
		}),
	}, protocol.MessageOptions{
		Event: &protocol.EventOptions{
			ID: &event.ID,
		},
	})

	PrintMessage(t, msg)
}

func Test_ValueAnalysisProducer(t *testing.T) {
	msg := p.NewValueAnalysisMessage(ts, protocol.SeverityInfo, []protocol.ValueAnalysis{
		protocol.NewAnalysis(protocol.Value(125123), "number is odd"),
	}, protocol.MessageOptions{
		Event: &protocol.EventOptions{
			ID: &event.ID,
		},
	})

	PrintMessage(t, msg)
}
