package pact

import "github.com/AsynkronIT/protoactor-go/actor"

func SpawnFromInstance(obj actor.Actor, prefix string, options ...Options) (*actor.PID, error) {
	return DEFAULT_PACT.SpawnFromInstance(obj, prefix, options...)
}

func SpawnFromFunc(f actor.ActorFunc, prefix string, options ...Options) (*actor.PID, error) {
	return DEFAULT_PACT.SpawnFromFunc(f, prefix, options...)
}
