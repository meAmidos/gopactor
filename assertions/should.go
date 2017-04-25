// Goconvey-style assertions to be used with actors
// spawned using Gopactor.
// It is most likely that you do not want to use this package directly,
// but rather import the higher level "github.com/meAmidos/gopactor".
package assertions

import "github.com/meamidos/gopactor/gopactor"

// ShouldReceive asserts that a given message is received by the actor
// and it does not matter who is the sender:
//   So(myActor, ShouldReceive, "ping")
func ShouldReceive(actual interface{}, expected ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldReceive(actual, expected...)
}

// ShouldReceiveFrom asserts that a given message is received by the actor
// and it is received from a certain sender:
//   So(myActor, ShouldReceiveFrom, sender, "ping")
func ShouldReceiveFrom(actual interface{}, expected ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldReceiveFrom(actual, expected...)
}

// ShouldReceiveSomething asserts that some message is received by the actor
// and it does not matter who is the sender
// and what is in the message:
//   So(myActor, ShouldReceiveSomething)
func ShouldReceiveSomething(actual interface{}, expected ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldReceiveSomething(actual, expected...)
}

// ShouldReceiveN asserts that any N messages are received by the actor
// and it does not matter who has sent them
// and what is the content of the messages:
//   So(myActor, ShouldReceiveN)
func ShouldReceiveN(actual interface{}, params ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldReceiveN(actual, params...)
}

// ShouldSend asserts that a given message is sent by the actor
// and it does not matter who is the receiver:
//   So(myActor, ShouldSend, "ping")
func ShouldSend(actual interface{}, expected ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldSend(actual, expected...)
}

// ShouldSendTo asserts that a given message is sent by the actor
// and it is addressed to a certain receiver:
//   So(myActor, ShouldSendTo, receiver, "ping")
func ShouldSendTo(actual interface{}, expected ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldSendTo(actual, expected...)
}

// ShouldSendSomething asserts that some message is sent by the actor
// and it does not matter who is the receiver
// and what is in the message:
//   So(myActor, ShouldSendSomething)
func ShouldSendSomething(actual interface{}, _ ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldSendSomething(actual)
}

// ShouldSendN asserts that any N messages are sent by the actor
// and it does not matter who they are addressed to
// and what is the content of the messages:
//   So(myActor, ShouldSendN)
func ShouldSendN(actual interface{}, params ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldSendN(actual, params...)
}

// ShouldNotSendOrReceive asserts that the actor does not send or receive
// anything during the given period of time (which you specify
// in options when you spawn the actor using Gopactor).
//   So(myActor, ShouldNotSendOrReceive)
func ShouldNotSendOrReceive(actual interface{}, _ ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldNotSendOrReceive(actual)
}

// ShouldStart asserts that the actor has formally started.
// That is, it has received the &actor.Started{} message.
//   So(myActor, ShouldStart)
func ShouldStart(actual interface{}, _ ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldStart(actual)
}

// ShouldStop asserts that the actor has stopped.
//   So(myActor, ShouldStop)
func ShouldStop(actual interface{}, _ ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldStop(actual)
}

// ShouldBeRestarting asserts that the actor is restarting
//   So(myActor, ShouldBeRestarting)
func ShouldBeRestarting(actual interface{}, _ ...interface{}) string {
	return gopactor.DEFAULT_GOPACTOR.ShouldBeRestarting(actual)
}
