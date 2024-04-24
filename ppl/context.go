package ppl

import "context"

type Context struct {
	c context.Context
}

func NewContext(initial context.Context) *Context {
	return &Context{
		c: initial,
	}
}

func (c *Context) Set(key any, value any) {
	c.c = context.WithValue(c.c, key, value)
}

func (c Context) Get(key any) any {
	return c.c.Value(key)
}

func (c *Context) Inner() context.Context {
	return c.c
}
