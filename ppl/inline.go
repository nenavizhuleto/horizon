package ppl

type inline struct {
	f    InlineElmFunc
	next Elm
}

func (h *inline) Next(next Elm) {
	h.next = next
}

func (h *inline) Message(ctx *Context, message any) {
	if h.f(ctx, message) && h.next != nil {
		h.next.Message(ctx, message)
	}
}

func Inline(f InlineElmFunc) Elm {
	return &inline{
		f: f,
	}
}
