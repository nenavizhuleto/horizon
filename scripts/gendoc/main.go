package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/nenavizhuleto/horizon/producer"
	"github.com/nenavizhuleto/horizon/protocol"
)

// Static example variables
var (
	// Output
	out = os.Stdout

	ts = time.Now()                                                     // timestamp
	p  = producer.New("PRODUCER_ID", "PRODUCER_NAME", "PRODUCER_GROUP") // producer

	// frame location
	loc = protocol.FrameLocation{
		Partition: 124,
		Offset:    5123512,
		Topic:     "TOPIC_NAME",
	}

	// event
	event = protocol.Event{
		ID:       "EVENT_ID",
		Start:    ts,
		End:      ts,
		Duration: 10 * time.Second,
	}

	// bounding box
	bbox = protocol.Position{
		X:      123,
		Y:      321,
		Width:  400,
		Height: 600,
	}

	// plate detections
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

	// object detections
	objects = []protocol.ObjectDetection{
		{
			ID:          "e422664a-d910-4a75-905c-1fdff28cd876",
			Class:       "person",
			BoundingBox: bbox,
			Confidence:  0.51,
		},
		{
			ID:          "ff74a087-5878-474c-9990-fb08b195a506",
			Class:       "car",
			BoundingBox: bbox,
			Confidence:  0.87,
		},
	}

	frame      = protocol.Frame{0x0, 0x1, 0x2, 0x3}
	dimensions = protocol.Dimensions{
		Width:  4123,
		Height: 1523,
	}
	positions = []protocol.Position{
		{X: 10, Y: 20, Width: 400, Height: 300},
		{X: 20, Y: 10, Width: 300, Height: 400},
	}

	// motion detections
	motions = []protocol.MotionDetection{
		{
			Source: protocol.Source{
				URI:        "http://localhost:1234/hello_world",
				Dimensions: dimensions,
			},
			Position: bbox,
		},
	}

	recording = protocol.Recording{
		File: "recording.mp4",
		URL:  "http://localhost:8912/recording.mp4",
	}

	media = []protocol.Media{
		{
			AnalysisID: "ANALYSIS_ID",
			Timestamp:  ts,
			File:       "object.jpeg",
			URL:        "http://localhost:8912/object.jpeg",
		},
	}

	vehicles = []protocol.Vehicle{
		{
			ID:    "VEHICLE_ID",
			Plate: "EM322",
			Brand: "Hyundai",
			Model: "Solaris",
			Color: "white",
		},
	}

	empty_message = protocol.Message[any]{
		Producer: protocol.Producer(p),
		Type:     protocol.MessageObjectDetection,
		Body:     "BODY",
	}
)

var entities = []any{
	empty_message,
	positions,
	loc,
	bbox,
	frame,
	dimensions,
	recording,
	media,
	event,
	motions,
	objects,
	plates,
	vehicles,
}

// Messages
var messages = []any{
	p.NewMotionDetectionMessage(ts, motions),
	p.NewObjectDetectionMessage(ts, loc, objects),
	p.NewObjectAnalysisMessage(event.ID, ts, protocol.SeverityInfo, loc, []protocol.ObjectAnalysis{
		{
			ID:      "OBJECT_ANALYSIS_ID",
			Subject: objects[0],
			Report:  "OBJECT_ANALYSIS_REPORT",
		},
	}),
	p.NewPlateDetectionMessage(ts, loc, plates),
	p.NewPlateAnalysisMessage(event.ID, ts, protocol.SeverityWarning, loc, []protocol.PlateAnalysis{
		{
			ID:      "PLATE_ANALYSIS_ID",
			Subject: plates[0],
			Report: protocol.PlateReport{
				{
					ID:       "LIST_ID",
					Name:     "LIST_NAME",
					Severity: protocol.SeverityWarning,
					Color:    "#00ff44",
					Vehicles: vehicles,
				},
			},
		},
	}),
	p.NewFrameMessage(ts, frame, dimensions, positions),
	p.NewEventStartMessage(event.ID, event.Start),
	p.NewEventEndMessage(event.ID, event.End, event.Duration),
	p.NewMediaMessage(event.ID, recording, media),
}

var (
	CRLF = "\r\n"

	MESSAGE_H      = "# Messages\r\n"
	MESSAGE_TYPE_T = "## Message type: `%s`\r\n\r\n"

	ENTITY_H      = "# Entities\r\n"
	ENTITY_TYPE_T = "## Entity type: `%s`\r\n\r\n"

	EXAMPLE_S   = "Example:\r\n"
	CODEBLOCK_T = "```go\r\n%s\r\n```\r\n"

	JSON_INDENT = "  "
)

func main() {

	out.WriteString(CRLF)
	out.WriteString(MESSAGE_H)
	for _, m := range messages {
		// Write template
		out.WriteString(CRLF)
		t := reflect.ValueOf(m).FieldByName("Type")
		out.WriteString(fmt.Sprintf(MESSAGE_TYPE_T, t))

		// Write message
		if bytes, err := json.MarshalIndent(m, "", JSON_INDENT); err != nil {
			os.Stderr.WriteString(fmt.Sprintf("failed to marshal message(%s): %v\n", t, err))
		} else {
			out.WriteString(fmt.Sprintf(CODEBLOCK_T, string(bytes)))
		}
	}

	out.WriteString(CRLF)
	out.WriteString(ENTITY_H)
	for _, e := range entities {
		out.WriteString(CRLF)
		t := reflect.ValueOf(e).Type().String()
		out.WriteString(fmt.Sprintf(ENTITY_TYPE_T, t))

		// Write message
		if bytes, err := json.MarshalIndent(e, "", JSON_INDENT); err != nil {
			os.Stderr.WriteString(fmt.Sprintf("failed to marshal entity(%s): %v\n", t, err))
		} else {
			out.WriteString(fmt.Sprintf(CODEBLOCK_T, string(bytes)))
		}
	}
}
