package pact

import (
	"testing"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/stretchr/testify/assert"
)

type TestActor struct{}

func (ta *TestActor) Receive(ctx actor.Context) {}

func TestSpawnFromInstance(t *testing.T) {
	a := assert.New(t)

	ta := &TestActor{}
	_, err := SpawnFromInstance(ta, "rcv")
	a.Nil(err)

	// Cleanup
	PactReset()
}

func TestSpawnFromProducer(t *testing.T) {
	a := assert.New(t)

	f := func() actor.Actor {
		return &TestActor{}
	}

	_, err := SpawnFromProducer(f, "rcv")
	a.Nil(err)

	// Cleanup
	PactReset()
}

func TestSpawnFromFunc(t *testing.T) {
	a := assert.New(t)

	_, err := SpawnFromFunc(func(ctx actor.Context) {}, "rcv")
	a.Nil(err)

	// Cleanup
	PactReset()
}
