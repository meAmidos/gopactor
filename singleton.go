package pact

import "github.com/AsynkronIT/protoactor-go/actor"

func SpawnFromInstance(obj actor.Actor, prefix string) (*actor.PID, error) {
	return DEFAULT_PACT.SpawnFromInstance(obj, prefix)
}

func SpawnFromFunc(f actor.ActorFunc, prefix string) (*actor.PID, error) {
	return DEFAULT_PACT.SpawnFromFunc(f, prefix)
}
