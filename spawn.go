package pact

import "github.com/AsynkronIT/protoactor-go/actor"

func (p *Pact) spawn(props *actor.Props, prefix string) (*actor.PID, error) {
	catcher := NewCatcher()
	catcher.LoggingOn = p.LoggingOn

	pid, err := catcher.Spawn(props, prefix)
	if err != nil {
		return nil, err
	}

	p.CatchersByPID[pid.String()] = catcher

	return pid, nil
}

func (p *Pact) SpawnFromInstance(obj actor.Actor, prefix string) (*actor.PID, error) {
	props := actor.FromInstance(obj)
	return p.spawn(props, prefix)
}

func (p *Pact) SpawnFromFunc(f actor.ActorFunc, prefix string) (*actor.PID, error) {
	props := actor.FromFunc(f)
	return p.spawn(props, prefix)
}
