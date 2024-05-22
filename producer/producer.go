package producer

import (
	"encoding/json"
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

type Producer protocol.Producer

func New(id, name, group string) Producer {
	return Producer(protocol.Producer{
		ID:      id,
		Name:    name,
		Group:   group,
		Modules: make([]protocol.ModuleName, 0),
	})
}

func (p *Producer) SetModules(modules []protocol.ModuleName) {
	p.Modules = modules
}

type RawMessage = protocol.Message[json.RawMessage]
type RawDetectionMessage = protocol.Message[protocol.DetectionMessage[json.RawMessage]]
type RawAnalysesMessage = protocol.Message[protocol.AnalysesMessage[json.RawMessage]]

type MotionDetectionMessage = protocol.Message[protocol.MotionDetectionMessage]
type ObjectDetectionMessage = protocol.Message[protocol.ObjectDetectionMessage]
type ObjectAnalysesMessage = protocol.Message[protocol.ObjectAnalysesMessage]

type PlateDetectionMessage = protocol.Message[protocol.PlateDetectionMessage]
type PlateAnalysesMessage = protocol.Message[protocol.PlateAnalysesMessage]

type EventStartMessage = protocol.Message[protocol.EventStartMessage]
type EventEndMessage = protocol.Message[protocol.EventEndMessage]

type FrameMessage = protocol.Message[protocol.FrameMessage]
type MediaMessage = protocol.Message[protocol.MediaMessage]

func (p Producer) NewMotionDetectionMessage(ts time.Time, motions []protocol.MotionDetection) MotionDetectionMessage {
	return protocol.NewMotionDetectionMessage(protocol.Producer(p), protocol.MotionDetectionMessage{
		Timestamp:  ts,
		Detections: motions,
	})
}

func (p Producer) NewFrameMessage(ts time.Time, data protocol.Frame, dims protocol.Dimensions, regs []protocol.Position) FrameMessage {
	return protocol.NewFrameMessage(protocol.Producer(p), protocol.FrameMessage{
		Timestamp:  ts,
		Dimensions: dims,
		Regions:    regs,
		Data:       data,
	})
}

func (p Producer) NewObjectDetectionMessage(ts time.Time, loc protocol.FrameLocation, objects []protocol.ObjectDetection) ObjectDetectionMessage {
	return protocol.NewObjectDetectionMessage(protocol.Producer(p), protocol.ObjectDetectionMessage{
		Timestamp:     ts,
		Detections:    objects,
		FrameLocation: loc,
	})
}

func (p Producer) NewObjectAnalysesMessage(id, event_id string, ts time.Time, severity protocol.Severity, loc protocol.FrameLocation, analyses []protocol.ObjectAnalysis) ObjectAnalysesMessage {
	return protocol.NewObjectAnalysesMessage(protocol.Producer(p), protocol.ObjectAnalysesMessage{
		ID:            id,
		EventID:       event_id,
		Timestamp:     ts,
		Severity:      severity,
		FrameLocation: loc,
		Analyses:      analyses,
	})
}

func (p Producer) NewPlateDetectionMessage(ts time.Time, loc protocol.FrameLocation, plates []protocol.PlateDetection) PlateDetectionMessage {
	return protocol.NewPlateDetectionMessage(protocol.Producer(p), protocol.PlateDetectionMessage{
		Timestamp:     ts,
		Detections:    plates,
		FrameLocation: loc,
	})
}

func (p Producer) NewPlateAnalysesMessage(id, event_id string, ts time.Time, severity protocol.Severity, loc protocol.FrameLocation, analyses []protocol.PlateAnalysis) PlateAnalysesMessage {
	return protocol.NewPlateAnalysesMessage(protocol.Producer(p), protocol.PlateAnalysesMessage{
		ID:            id,
		EventID:       event_id,
		Timestamp:     ts,
		Severity:      severity,
		FrameLocation: loc,
		Analyses:      analyses,
	})
}

func (p Producer) NewEventStartMessage(id string, start time.Time) EventStartMessage {
	return protocol.NewEventStartMessage(protocol.Producer(p), protocol.EventStartMessage{
		ID:    id,
		Start: start,
	})
}

func (p Producer) NewEventEndMessage(id string, end time.Time, duration time.Duration) EventEndMessage {
	return protocol.NewEventEndMessage(protocol.Producer(p), protocol.EventEndMessage{
		ID:       id,
		End:      end,
		Duration: duration,
	})
}

func (p Producer) NewMediaMessage(event_id string, recording protocol.Recording, media []protocol.Media) MediaMessage {
	return protocol.NewMediaMessage(protocol.Producer(p), protocol.MediaMessage{
		EventID:   event_id,
		Recording: recording,
		Media:     media,
	})
}
