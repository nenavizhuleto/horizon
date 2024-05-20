package producer

import (
	"time"

	"github.com/nenavizhuleto/horizon/protocol"
)

type Producer protocol.Producer

func New(id, name, group string) Producer {
	return Producer(protocol.Producer{
		ID:      id,
		Name:    name,
		Group:   group,
		Modules: make([]string, 0),
	})
}

func (p *Producer) AddModule(module string) {
	p.Modules = append(p.Modules, module)
}

func (p *Producer) SetModules(modules []string) {
	p.Modules = modules
}

type MotionDetectionMessage = protocol.Message[protocol.MotionDetectionMessage]

func (p Producer) NewMotionDetectionMessage(ts time.Time, motions []protocol.MotionDetection) MotionDetectionMessage {
	return protocol.NewMotionDetectionMessage(protocol.Producer(p), protocol.MotionDetectionMessage{
		Timestamp:  ts,
		Detections: motions,
	})
}

type FrameMessage = protocol.Message[protocol.FrameMessage]

func (p Producer) NewFrameMessage(ts time.Time, data protocol.Frame, dims protocol.Dimensions, regs []protocol.Position) FrameMessage {
	return protocol.NewFrameMessage(protocol.Producer(p), protocol.FrameMessage{
		Timestamp:  ts,
		Dimensions: dims,
		Regions:    regs,
		Data:       data,
	})
}

type ObjectDetectionMessage = protocol.Message[protocol.ObjectDetectionMessage]

func (p Producer) NewObjectDetectionMessage(ts time.Time, loc protocol.Location, objects []protocol.ObjectDetection) ObjectDetectionMessage {
	return protocol.NewObjectDetectionMessage(protocol.Producer(p), protocol.ObjectDetectionMessage{
		Timestamp:  ts,
		Detections: objects,
		Location:   loc,
	})
}

type ObjectAnalysisMessage = protocol.Message[protocol.ObjectAnalysisMessage]

func (p Producer) NewObjectAnalysisMessage(event_id string, ts time.Time, severity protocol.Severity, loc protocol.Location, analyses []protocol.ObjectAnalysis) ObjectAnalysisMessage {
	return protocol.NewObjectAnalysisMessage(protocol.Producer(p), protocol.ObjectAnalysisMessage{
		EventID:   event_id,
		Timestamp: ts,
		Severity:  severity,
		Location:  loc,
		Analyses:  analyses,
	})
}

type PlateDetectionMessage = protocol.Message[protocol.PlateDetectionMessage]

func (p Producer) NewPlateDetectionMessage(ts time.Time, loc protocol.Location, plates []protocol.PlateDetection) PlateDetectionMessage {
	return protocol.NewPlateDetectionMessage(protocol.Producer(p), protocol.PlateDetectionMessage{
		Timestamp:  ts,
		Detections: plates,
		Location:   loc,
	})
}

type PlateAnalysisMessage = protocol.Message[protocol.PlateAnalysisMessage]

func (p Producer) NewPlateAnalysisMessage(event_id string, ts time.Time, severity protocol.Severity, loc protocol.Location, analyses []protocol.PlateAnalysis) PlateAnalysisMessage {
	return protocol.NewPlateAnalysisMessage(protocol.Producer(p), protocol.PlateAnalysisMessage{
		EventID:   event_id,
		Timestamp: ts,
		Severity:  severity,
		Location:  loc,
		Analyses:  analyses,
	})
}

type EventStartMessage = protocol.Message[protocol.EventStartMessage]

func (p Producer) NewEventStartMessage(id string, start time.Time) EventStartMessage {
	return protocol.NewEventStartMessage(protocol.Producer(p), protocol.EventStartMessage{
		ID:    id,
		Start: start,
	})
}

type EventEndMessage = protocol.Message[protocol.EventEndMessage]

func (p Producer) NewEventEndMessage(id string, end time.Time, duration time.Duration) EventEndMessage {
	return protocol.NewEventEndMessage(protocol.Producer(p), protocol.EventEndMessage{
		ID:       id,
		End:      end,
		Duration: duration,
	})
}

type MediaMessage = protocol.Message[protocol.MediaMessage]

func (p Producer) NewMediaMessage(event_id string, recording protocol.Recording, media []protocol.Media) MediaMessage {
	return protocol.NewMediaMessage(protocol.Producer(p), protocol.MediaMessage{
		EventID:   event_id,
		Recording: recording,
		Media:     media,
	})
}
