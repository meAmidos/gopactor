# Gopactor - testing for ProtoActor
Gopactor is a set of tools to simplify writing BDD tests for actors created with [Protoactor](https://github.com/AsynkronIT/protoactor-go).

[![GoDoc](https://godoc.org/github.com/meAmidos/gopactor?status.svg)](https://godoc.org/github.com/meAmidos/gopactor)

Currently, the main focus is to provide convenient assertions for tests written using the [Goconvey](http://goconvey.co/) framework. However, all provided assertions can potentially be used independently, and it is easy to write an adapter to whatever matcher/assertion library you prefer.

Any contribution to this project will be highly appreciated!

## Example of usage
Here is a short example. We'll define and test a simple worker actor that can do only one thing: respond "pong" when it receives "ping".

```
package worker_test

import (
    "testing"

    "github.com/AsynkronIT/protoactor-go/actor"
    . "github.com/meAmidos/gopactor"
    . "github.com/smartystreets/goconvey/convey"
)

// Actor to test
type Worker struct{}

// This actor can do only one thing, but it does this thing well.
func (w *Worker) Receive(ctx actor.Context) {
    switch m := ctx.Message().(type) {
    case string:
        if m == "ping" {
            ctx.Respond("pong")
        }
    }
}

// For the sake of simplicity, the test is flat and not structured.
// In real life, you may want to follow the "Given-When-Then" approach
// which is commonly recognized as a good BDD-style.
func TestWorker(t *testing.T) {
    Convey("Test the worker actor", t, func() {

        // It is essential to spawn the tested actor using Gopactor. This way, Gopactor
        // will be able to intercept all inbound/outbound messages of the actor.
        worker, err := SpawnFromInstance(&Worker{}, OptDefault.WithPrefix("worker"))
        So(err, ShouldBeNil)

        // Spawn an additional actor that will communicate with our worker.
        // The only purpose of this actor is to be a sparring partner,
        // so we don't care about its functionality.
        // Conveniently, Gopactor provides an easy way to create it.
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
```

## Main features
### Intercept messages
For any actor you want to test, Gopactor can intercept all it's inbound and outbound messages. It can be very useful when you want to test the actor's behavior. Moreover, interception forces a naturally asynchronous actor to act in a more synchronous way that is much easier to reason about. So, messages are sent and received under the control of Gopactor, step by step.

### Intercept system messages
Protoactor uses some specific system messages to control the lifecycle of an actor. Gopactor can intercept some of such messages to help you test that your actor stops or restarts when expected.

### Intercept spawning of children
It is a common pattern to let actors spawn child actors and communicate with them. Good as it is, this pattern often stays in the way of writing deterministic tests. Given that child-spawning and communication happen in the background asynchronously, it can be seen more like a side-effect that can interfere with our tests in many unpredictable ways.

So, the current Gopactor's approach is to intercept all spawn invocations and instead of spawning what is requested, spawn no-op null-actors that are guaranteed to not communicate with parents in any way. It is planned to evolve this approach to something even more useful and configurable in the future.

### Goconvey-style assertions
Gopactor provides a bunch of assertion functions to be used with the very popular testing framework Goconvey (http://goconvey.co/). For instance,

```
So(worker, ShouldReceive, "ping")
So(worker, ShouldSendTo, requestor, "pong")
```

### Configurable
For every tested actor, you can define what you want to intercept: inbound, outbound or system messages. Or everything. Or nothing at all. You can also set a custom timeout:

```
options := OptNoInterception.
    WithOutboundInterception().
    WithPrefix("my-actor").
    WithTimeout(10 * time.Millisecond)
```

## Supported assertions
```
ShouldReceive
ShouldReceiveFrom
ShouldReceiveSomething
ShouldReceiveN

ShouldSend
ShouldSendTo
ShouldSendSomething
ShouldSendN

ShouldNotSendOrReceive

ShouldStart
ShouldStop
ShouldBeRestarting
```

# Plans
Many things could be done to improve the library. Some of the areas that I am personally interested in (with no particular order):
- Review the interception of child-spawning
- Add assertions for spawning
- Catch more system messages
- Add an optional logger
- Add negative-scenario assertions (`ShouldNotReceive`, etc.)
- Be smart in handling/asserting actors failures
- Handle outbound system messages separately

# Contribution
Please feel free to open an issue if you encounter a problem with the library or have a question. Pull requests will be highly appreciated.
