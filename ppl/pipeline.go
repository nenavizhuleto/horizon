package ppl

type pipeline struct {
	start Elm
	end   Elm
}

func (p *pipeline) Message(ctx *Context, message any) {
	if p.start != nil {
		p.start.Message(ctx, message)
	}
}

func (p *pipeline) Next(next Elm) {
	if end, ok := p.end.(nextable); ok {
		end.Next(next)
	}
}

func NewPipeline(elms ...Elm) Elm {
	if len(elms) == 0 {
		return &pipeline{}
	}

	if len(elms) == 1 {
		return &pipeline{
			start: elms[0],
		}
	}

	for i := 1; i < len(elms); i++ {
		if elm, ok := elms[i-1].(nextable); ok {
			elm.Next(elms[i])
		}
	}

	return &pipeline{
		start: elms[0],
		end:   elms[len(elms)-1],
	}
}
