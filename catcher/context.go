package catcher

import "github.com/AsynkronIT/protoactor-go/actor"

type NullReceiver struct{}

func (nr *NullReceiver) Receive(ctx actor.Context) {}

// This is a wrapper around a real actor.Context object
// to intercept calls for testing purposes.
// This should implement the actor.Context interface
type Context struct {
	catcher       *Catcher
	actor.Context // This is the original context to pass calls to
}

func NewContext(catcher *Catcher, ctx actor.Context) *Context {
	return &Context{catcher, ctx}
}

// Intercept Spawn calls to create dummy actors instead of real ones
func (ctx *Context) Spawn(_ *actor.Props) *actor.PID {
	props := actor.FromInstance(&NullReceiver{})
	return ctx.Context.Spawn(props)
}

func (ctx *Context) SpawnPrefix(_ *actor.Props, prefix string) *actor.PID {
	props := actor.FromInstance(&NullReceiver{})
	return ctx.Context.SpawnPrefix(props, prefix)
}

func (ctx *Context) SpawnNamed(_ *actor.Props, id string) (*actor.PID, error) {
	props := actor.FromInstance(&NullReceiver{})
	return ctx.Context.SpawnNamed(props, id)
}
