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
	_, err := SpawnFromInstance(ta)
	a.Nil(err)

	// Cleanup
	PactReset()
}

func TestSpawnFromInstance_WithPrefix(t *testing.T) {
	a := assert.New(t)

	ta := &TestActor{}
	object, err := SpawnFromInstance(ta, OptDefault.WithPrefix("test-actor"))
	a.Nil(err)
	a.Contains(object.String(), "test-actor")

	// Cleanup
	PactReset()
}

func TestSpawnFromProducer(t *testing.T) {
	a := assert.New(t)

	f := func() actor.Actor {
		return &TestActor{}
	}

	_, err := SpawnFromProducer(f)
	a.Nil(err)

	// Cleanup
	PactReset()
}

func TestSpawnFromFunc(t *testing.T) {
	a := assert.New(t)

	_, err := SpawnFromFunc(func(ctx actor.Context) {})
	a.Nil(err)

	// Cleanup
	PactReset()
}

func TestSpawnMock(t *testing.T) {
	a := assert.New(t)

	_, err := SpawnMock()
	a.Nil(err)

	// Cleanup
	PactReset()
}
