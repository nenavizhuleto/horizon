package ops

import (
	"github.com/nenavizhuleto/horizon/ppl"
	"github.com/nenavizhuleto/horizon/protocol"
)

type DFilter interface {
	ppl.Elm
	Motions(elms ...ppl.Elm) DFilter
	Objects(elms ...ppl.Elm) DFilter
	Values(elms ...ppl.Elm) DFilter
}

type detection_filter struct {
	motions ppl.Elm
	objects ppl.Elm
	values  ppl.Elm

	next ppl.Elm
}

func DetectionFilter() DFilter {
	return &detection_filter{}
}

func (h *detection_filter) Next(next ppl.Elm) {
	h.next = next
}

func (h *detection_filter) Motions(elms ...ppl.Elm) DFilter {
	h.motions = ppl.NewPipeline(elms...)
	return h
}

func (h *detection_filter) Objects(elms ...ppl.Elm) DFilter {
	h.objects = ppl.NewPipeline(elms...)
	return h
}

func (h *detection_filter) Values(elms ...ppl.Elm) DFilter {
	h.values = ppl.NewPipeline(elms...)
	return h
}

func (h *detection_filter) Message(ctx *ppl.Context, message any) {
	switch detection := message.(type) {
	case protocol.MotionDetectionMessage:
		if h.motions == nil {
			break
		}
		for _, motion := range detection.Motions {
			h.motions.Message(ctx, motion)
		}
	case protocol.ObjectDetectionMessage:
		if h.objects == nil {
			break
		}
		for _, object := range detection.Objects {
			h.objects.Message(ctx, object)
		}
	case protocol.ObjectDetectionMessage:
		if h.objects == nil {
			break
		}
		for _, object := range detection.Objects {
			h.objects.Message(ctx, object)
		}
	case protocol.ValueDetectionMessage:
		if h.values == nil {
			break
		}
		for _, value := range detection.Values {
			h.values.Message(ctx, value)
		}
	}

	if h.next != nil {
		h.next.Message(ctx, message)
	}
}
