package gopactor

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"gopactor/gopactor"
	"gopactor/options"
)

// Analog of Protoactor's actor.SpawnPrefix(actor.FromInstance(...))
// The main difference is that after spawning with Gopactor
// you can write assertions for the spawned actor.
func SpawnFromInstance(obj actor.Actor, opts ...options.Options) (*actor.PID, error) {
	return gopactor.DEFAULT_GOPACTOR.SpawnFromInstance(obj, opts...)
}

// Analog of Protoactor's actor.SpawnPrefix(actor.FromProducer(...))
func SpawnFromProducer(producer actor.Producer, opts ...options.Options) (*actor.PID, error) {
	return gopactor.DEFAULT_GOPACTOR.SpawnFromProducer(producer, opts...)
}

// Analog of Protoactor's actor.SpawnPrefix(actor.FromFunc(...))
func SpawnFromFunc(f actor.ActorFunc, opts ...options.Options) (*actor.PID, error) {
	return gopactor.DEFAULT_GOPACTOR.SpawnFromFunc(f, opts...)
}

// Spawn an actor that does nothing.
// It can be very useful in tests when all you need is an actor
// that can play a role of a message sender and a black-hole receiver.
func SpawnNullActor(opts ...options.Options) (*actor.PID, error) {
	return gopactor.DEFAULT_GOPACTOR.SpawnNullActor(opts...)
}

// PactReset cleans up internal data structures used by Gopactor.
// Normally, you do not have to use it. If you just test a dozen of actors
// in a short-living test, there is no need to care about cleaning up.
// However, if for some reason, you are spawning thousands of actors in a long-running
// test, you might want to call this function from time to time.
func PactReset() {
	gopactor.DEFAULT_GOPACTOR.Reset()
}
