package pact

import "github.com/AsynkronIT/protoactor-go/actor"

func (p *Pact) spawn(props *actor.Props, prefix string, options ...Options) (*actor.PID, error) {
	catcher := NewCatcher()
	catcher.LoggingOn = p.LoggingOn

	pid, err := catcher.Spawn(props, prefix, options...)
	if err != nil {
		return nil, err
	}

	p.CatchersByPID[pid.String()] = catcher

	return pid, nil
}

func (p *Pact) SpawnFromInstance(obj actor.Actor, prefix string, options ...Options) (*actor.PID, error) {
	props := actor.FromInstance(obj)
	return p.spawn(props, prefix, options...)
}

func (p *Pact) SpawnFromFunc(f actor.ActorFunc, prefix string, options ...Options) (*actor.PID, error) {
	props := actor.FromFunc(f)
	return p.spawn(props, prefix, options...)
}
