package gopactor

import (
	"github.com/AsynkronIT/protoactor-go/actor"

	"gopactor/catcher"
	"gopactor/options"
)

func (p *Gopactor) spawn(props *actor.Props, opts ...options.Options) (*actor.PID, error) {
	catcher := catcher.New()

	pid, err := catcher.Spawn(props, opts...)
	if err != nil {
		return nil, err
	}

	p.CatchersByPID[pid.String()] = catcher

	return pid, nil
}

func (p *Gopactor) SpawnFromInstance(obj actor.Actor, opts ...options.Options) (*actor.PID, error) {
	props := actor.FromInstance(obj)
	return p.spawn(props, opts...)
}

func (p *Gopactor) SpawnFromProducer(producer actor.Producer, opts ...options.Options) (*actor.PID, error) {
	props := actor.FromProducer(producer)
	return p.spawn(props, opts...)
}

func (p *Gopactor) SpawnFromFunc(f actor.ActorFunc, opts ...options.Options) (*actor.PID, error) {
	props := actor.FromFunc(f)
	return p.spawn(props, opts...)
}

func (p *Gopactor) SpawnNullActor(opts ...options.Options) (*actor.PID, error) {
	return p.SpawnFromInstance(&catcher.NullReceiver{}, opts...)
}
