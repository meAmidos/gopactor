package pact

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/meamidos/pact/catcher"
	"github.com/meamidos/pact/pact"
)

func SpawnFromInstance(obj actor.Actor, prefix string, options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnFromInstance(obj, prefix, options...)
}

func SpawnFromFunc(f actor.ActorFunc, prefix string, options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnFromFunc(f, prefix, options...)
}

func SpawnMockWithPrefix(prefix string, options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnMockWithPrefix(prefix, options...)
}

func PactReset() {
	pact.DEFAULT_PACT.Reset()
}
