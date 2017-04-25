package pact

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/meamidos/pact/catcher"
	"github.com/meamidos/pact/options"
)

func (p *Pact) spawn(props *actor.Props, opts ...options.Options) (*actor.PID, error) {
	catcher := catcher.New()
	catcher.LoggingOn = p.LoggingOn

	pid, err := catcher.Spawn(props, opts...)
	if err != nil {
		return nil, err
	}

	p.CatchersByPID[pid.String()] = catcher

	return pid, nil
}

func (p *Pact) SpawnFromInstance(obj actor.Actor, opts ...options.Options) (*actor.PID, error) {
	props := actor.FromInstance(obj)
	return p.spawn(props, opts...)
}

func (p *Pact) SpawnFromProducer(producer actor.Producer, opts ...options.Options) (*actor.PID, error) {
	props := actor.FromProducer(producer)
	return p.spawn(props, opts...)
}

func (p *Pact) SpawnFromFunc(f actor.ActorFunc, opts ...options.Options) (*actor.PID, error) {
	props := actor.FromFunc(f)
	return p.spawn(props, opts...)
}

func (p *Pact) SpawnNullActor(opts ...options.Options) (*actor.PID, error) {
	return p.SpawnFromInstance(&catcher.NullReceiver{}, opts...)
}
