/*
Package Pact provides a set of tools to simplify testing
of actors created with Protoactor (https://github.com/AsynkronIT/protoactor-go).

Main features:

Intercept messages

For any actor you want to test, Pact can intercept all it's inbound and outbound
messages. It can be very useful when you want to test the actor's behavior. Moreover,
interception forces a naturally asynchronous actor to act in a more synchronous way that
is much easier to reason about. So, messages are sent and received under the
control of Pact, step by step.

Intercept system messages

Protoactor uses some specific system messages to control the lifecycle of an actor.
Pact can intercept some of such messages to help you test that your actor stops or restarts
when expected.

Intercept spawning of children

It is a common pattern to let actors spawn child actors and communicate with them.
Good as it is, this pattern often stays in the way of writing deterministic tests.
Given that child-spawning and communication happen in the background asynchronously,
it can be seen more like a side-effect that can interfere with our tests in many
unpredictable ways.

So, the current Pact's approach is to intercept all spawn invocations and instead of
spawning what is requested, spawn no-op null-actors that are guaranteed to not communicate
with parents in any way. It is planned to evolve this approach to something even more
useful and configurable in the future.

Goconvey-style assertions

Pact provides a bunch of assertion functions to be used with the very popular testing
framework Goconvey (http://goconvey.co/). For instance,

	So(worker, ShouldReceive, "ping")
	So(worker, ShouldSendTo, requestor, "pong")

Configurable

For every tested actor, you can define what you want to intercept: inbound, outbound
or system messages. Or everything. Or nothing at all. You can also set a custom timeout:

	options := OptNoInterception.
		WithOutboundInterception().
		WithPrefix("my-actor").
		WithTimeout(10 * time.Millisecond)

Example of usage

Here is a short example. We'll define and test a simple worker actor that can do only one thing:
respond "pong" when it receives "ping".

	package worker_test

	import (
		"testing"

		"github.com/AsynkronIT/protoactor-go/actor"
		. "github.com/meAmidos/pact"
		. "github.com/smartystreets/goconvey/convey"
	)

	// Actor to test
	type Worker struct{}

	// This actor is very simple. It can do only one thing, but it does this thing well.
	func (w *Worker) Receive(ctx actor.Context) {
		switch m := ctx.Message().(type) {
		case string:
			if m == "ping" {
				ctx.Respond("pong")
			}
		}
	}

	func TestWorker(t *testing.T) {
		Convey("Test the worker actor", t, func() {

			// It is essential to spawn the tested actor using Pact. This way, Pact
			// will be able to intercept all inbound/outbound messages of the actor.
			worker, err := SpawnFromInstance(&Worker{}, OptDefault.WithPrefix("worker"))
			So(err, ShouldBeNil)

			// Spawn an additional actor that will communicate with our worker.
			// The only purpose of this actor is to be a sparring partner,
			// so we don't care about its functionality.
			// Conveniently, Pact provides an easy way to create it.
			requestor, err := SpawnNullActor()
			So(err, ShouldBeNil)

			// Let the requestor ping the worker
			worker.Request("ping", requestor)

			// Assert that the worker receives the ping message
			So(worker, ShouldReceive, "ping")

			// Assert that the worker sends back the correct response
			So(worker, ShouldSendTo, requestor, "pong")

			// Finally, assert that the requestor gets the response
			So(requestor, ShouldReceive, "pong")
		})
	}

*/
package pact

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/meamidos/pact/catcher"
	"github.com/meamidos/pact/pact"
)

// Analog of Protoactor's actor.SpawnPrefix(actor.FromInstance(...))
// The main difference is that after spawning with Pact
// you can write assertions for the spawned actor.
func SpawnFromInstance(obj actor.Actor, options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnFromInstance(obj, options...)
}

// Analog of Protoactor's actor.SpawnPrefix(actor.FromProducer(...))
func SpawnFromProducer(producer actor.Producer, options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnFromProducer(producer, options...)
}

// Analog of Protoactor's actor.SpawnPrefix(actor.FromFunc(...))
func SpawnFromFunc(f actor.ActorFunc, options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnFromFunc(f, options...)
}

// Spawn an actor that does nothing.
// It can be very useful in tests when all you need is an actor
// that can play a role of a message sender and a black-hole receiver.
func SpawnNullActor(options ...catcher.Options) (*actor.PID, error) {
	return pact.DEFAULT_PACT.SpawnNullActor(options...)
}

// PactReset cleans up internal data structures used by Pact.
// Normally, you do not have to use it. If you just test a dozen of actors
// in a short-living test, there is no need to care about cleaning up.
// However, if for some reason, you are spawning thousands of actors in a long-running
// test, you might want to call this function from time to time.
func PactReset() {
	pact.DEFAULT_PACT.Reset()
}
