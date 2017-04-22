package pact

import "github.com/AsynkronIT/protoactor-go/actor"

var DEFAULT_PACT *Pact

func init() {
	DEFAULT_PACT = New()
}

func SpawnFromInstance(obj actor.Actor, prefix string, options ...Options) (*actor.PID, error) {
	return DEFAULT_PACT.SpawnFromInstance(obj, prefix, options...)
}

func SpawnFromFunc(f actor.ActorFunc, prefix string, options ...Options) (*actor.PID, error) {
	return DEFAULT_PACT.SpawnFromFunc(f, prefix, options...)
}

func PactReset() {
	DEFAULT_PACT.Reset()
}
