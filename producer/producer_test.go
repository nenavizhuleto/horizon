package producer_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/nenavizhuleto/horizon/producer"
	"github.com/nenavizhuleto/horizon/protocol"
)

var (
	ts  = time.Now()
	p   = producer.New("PRODUCER_ID", "PRODUCER_NAME", "PRODUCER_GROUP")
	loc = protocol.FrameLocation{
		Partition: 124,
		Offset:    5123512,
		Topic:     "TOPIC_NAME",
	}

	event = protocol.Event{
		ID:       "EVENT_ID",
		Start:    ts,
		End:      ts,
		Duration: 10 * time.Second,
	}

	bbox = protocol.Position{
		X:      123,
		Y:      321,
		Width:  400,
		Height: 600,
	}

	plates = []protocol.PlateDetection{
		{
			BoundingBox: bbox,
			Confidence:  0.99,
			Plate:       "EM322",
		},
		{
			BoundingBox: bbox,
			Confidence:  0.64,
			Plate:       "AD231",
		},
	}

	objects = []protocol.ObjectDetection{
		{
			ID:          "e422664a-d910-4a75-905c-1fdff28cd876",
			Class:       "car",
			BoundingBox: bbox,
			Confidence:  0.51,
		},
		{
			ID:          "ff74a087-5878-474c-9990-fb08b195a506",
			Class:       "chair",
			BoundingBox: bbox,
			Confidence:  0.87,
		},
	}

	motions = []protocol.MotionDetection{
		{
			Source: protocol.Source{
				URI: "http://localhost:1234/hello_world",
				Dimensions: protocol.Dimensions{
					Width:  4123,
					Height: 1523,
				},
			},
			Position: bbox,
		},
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
	msg := p.NewMotionDetectionMessage(ts, motions)
	PrintMessage(t, msg)
}

func Test_ObjectDetectionProducer(t *testing.T) {
	msg := p.NewObjectDetectionMessage(ts, loc, objects)
	PrintMessage(t, msg)
}

func Test_ObjectAnalysisProducer(t *testing.T) {
	msg := p.NewObjectAnalysesMessage("OBJECT_ANALYSES_ID", event.ID, ts, protocol.SeverityInfo, loc, []protocol.ObjectAnalysis{
		{
			Subject: objects[0],
			Report:  "OBJECT_ANALYSIS_REPORT",
		},
	})

	PrintMessage(t, msg)
}

func Test_PlateDetectionProducer(t *testing.T) {
	msg := p.NewPlateDetectionMessage(ts, loc, plates)
	PrintMessage(t, msg)
}

func Test_PlateAnalysisProducer(t *testing.T) {
	msg := p.NewPlateAnalysesMessage("PLATE_ANALYSES_ID", event.ID, ts, protocol.SeverityWarning, loc, []protocol.PlateAnalysis{
		{
			Subject: plates[0],
			Report: protocol.PlateReport{
				{
					ID:       "LIST_ID",
					Name:     "LIST_NAME",
					Severity: protocol.SeverityWarning,
					Color:    "#00ff44",
				},
			},
		},
	})
	PrintMessage(t, msg)
}

func Test_FrameProducer(t *testing.T) {
	msg := p.NewFrameMessage(ts, protocol.Frame{}, protocol.Dimensions{}, []protocol.Position{})
	PrintMessage(t, msg)
}

func Test_EventProducer(t *testing.T) {
	start_msg := p.NewEventStartMessage(event.ID, event.Start)
	end_msg := p.NewEventEndMessage(event.ID, event.End, event.Duration)

	PrintMessage(t, start_msg)
	PrintMessage(t, end_msg)
}

func Test_MediaProducer(t *testing.T) {
	msg := p.NewMediaMessage(event.ID, protocol.Recording{
		File: "recording.mp4",
		URL:  "http://localhost:8912/recording.mp4",
	}, []protocol.Media{
		{
			AnalysisID: "ANALYSIS_ID",
			Timestamp:  ts,
			File:       "object.jpeg",
			URL:        "http://localhost:8912/object.jpeg",
		},
	})

	PrintMessage(t, msg)
}
