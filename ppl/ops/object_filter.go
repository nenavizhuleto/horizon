package ops

import (
	"github.com/nenavizhuleto/horizon/ppl"
	"github.com/nenavizhuleto/horizon/protocol"
)

type object_elm struct {
	f    ObjectFunc
	next ppl.Elm
}

func ObjectElm(f ObjectFunc) ppl.Elm {
	return &object_elm{
		f: f,
	}
}

func (h *object_elm) Next(next ppl.Elm) {
	h.next = next
}

func (h *object_elm) Message(ctx *ppl.Context, message any) {
	if object, ok := message.(protocol.Object); ok {
		h.f(ctx, object)
	}

	if h.next != nil {
		h.next.Message(ctx, message)
	}
}

type object_filter struct {
	obj_elm ppl.Elm
	class   string

	next ppl.Elm
}

type ObjectFunc func(ctx *ppl.Context, object protocol.Object)

func ObjectFilter(class string, f ObjectFunc) ppl.Elm {
	return &object_filter{
		obj_elm: ObjectElm(f),
		class:   class,
	}
}

func (f *object_filter) Next(next ppl.Elm) {
	f.next = next
}

func (f *object_filter) Message(ctx *ppl.Context, message any) {
	if object, ok := message.(protocol.Object); ok {
		if object.Class == f.class {
			f.obj_elm.Message(ctx, object)
		}
	}

	if f.next != nil {
		f.next.Message(ctx, message)
	}
}
