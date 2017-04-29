package catcher

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

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

func (ctx *Context) Spawn(props *actor.Props) *actor.PID {
	catcher := ctx.catcher
	if catcher.Options.DummySpawningEnabled {
		props = actor.FromInstance(&NullReceiver{})
	}

	pid := ctx.Context.Spawn(props)
	if catcher.Options.SpawnInterceptionEnabled {
		catcher.ChSpawning <- pid
	}

	return pid
}

func (ctx *Context) SpawnPrefix(props *actor.Props, prefix string) *actor.PID {
	catcher := ctx.catcher
	if catcher.Options.DummySpawningEnabled {
		props = actor.FromInstance(&NullReceiver{})
	}

	pid := ctx.Context.SpawnPrefix(props, prefix)
	if catcher.Options.SpawnInterceptionEnabled {
		catcher.ChSpawning <- pid
	}

	return pid
}

func (ctx *Context) SpawnNamed(props *actor.Props, id string) (*actor.PID, error) {
	catcher := ctx.catcher
	if catcher.Options.DummySpawningEnabled {
		props = actor.FromInstance(&NullReceiver{})
	}

	pid, err := ctx.Context.SpawnNamed(props, id)
	if err == nil && catcher.Options.SpawnInterceptionEnabled {
		catcher.ChSpawning <- pid
	}

	return pid, err
}
