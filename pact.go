package pact

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/meamidos/pact/catcher"
	"github.com/meamidos/pact/pact"
)

func SpawnFromInstance(obj actor.Actor, options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnFromInstance(obj, options...)
}

func SpawnFromProducer(producer actor.Producer, options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnFromProducer(producer, options...)
}

func SpawnFromFunc(f actor.ActorFunc, options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnFromFunc(f, options...)
}

func SpawnMock(options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnMock(options...)
}

func PactReset() {
	pact.DEFAULT_PACT.Reset()
}
