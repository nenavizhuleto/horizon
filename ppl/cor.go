package ppl

type ElmFunc func(ctx *Context, message any)

type InlineElmFunc func(ctx *Context, message any) bool

type Elm interface {
	Message(ctx *Context, message any)
}

type nextable interface {
	Next(Elm)
}

type elm struct {
	f    ElmFunc
	next Elm
}

func NewElm(f ElmFunc) Elm {
	return &elm{
		f: f,
	}
}

func (e *elm) Next(next Elm) {
	e.next = next
}

func (e *elm) Message(ctx *Context, message any) {
	e.f(ctx, message)

	if e.next != nil {
		e.next.Message(ctx, message)
	}
}
