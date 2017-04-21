package pact

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

//
// Singleton interface
//

func ShouldReceive(actual interface{}, expected ...interface{}) string {
	return DEFAULT_PACT.ShouldReceive(actual, expected...)
}

func ShouldReceiveFrom(actual interface{}, expected ...interface{}) string {
	return DEFAULT_PACT.ShouldReceiveFrom(actual, expected...)
}

func ShouldReceiveN(actual interface{}, params ...interface{}) string {
	return DEFAULT_PACT.ShouldReceiveN(actual, params...)
}

func ShouldStop(actual interface{}, _ ...interface{}) string {
	return DEFAULT_PACT.ShouldStop(actual)
}

func ShouldSend(actual interface{}, expected ...interface{}) string {
	return DEFAULT_PACT.ShouldSend(actual, expected...)
}

func ShouldSendTo(actual interface{}, expected ...interface{}) string {
	return DEFAULT_PACT.ShouldSendTo(actual, expected...)
}

func ShouldSendSomething(actual interface{}, _ ...interface{}) string {
	return DEFAULT_PACT.ShouldSendSomething(actual)
}

func ShouldSendN(actual interface{}, params ...interface{}) string {
	return DEFAULT_PACT.ShouldSendSomething(actual, params)
}

func ShouldNotSendOrReceive(actual interface{}, _ ...interface{}) string {
	return DEFAULT_PACT.ShouldSendSomething(actual)
}

//
// Object interface
//

// Should receive a given message.
// It does not matter who is the sender.
func (p *Pact) ShouldReceive(param1 interface{}, params ...interface{}) string {
	actual, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	if len(params) != 1 {
		return "One parameter with a message is required to assert receiving"
	}

	expectedMsg, ok := params[0].(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	return p.shouldReceive(actual, nil, expectedMsg)
}

// Should receive a given message from a given sender
func (p *Pact) ShouldReceiveFrom(param1 interface{}, params ...interface{}) string {
	actual, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	if len(params) != 2 {
		return "Two parameters are required to assert receiving"
	}

	expectedMsg := params[0]

	// Two arguments means that the second is the expected sender
	sender, ok := params[1].(*actor.PID)
	if !ok {
		return "Sender should be an actor PID"
	}

	return p.shouldReceive(actual, sender, expectedMsg)
}

// Should receive N any messages
func (p *Pact) ShouldReceiveN(param1 interface{}, params ...interface{}) string {
	actual, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	if len(params) != 1 {
		return "One paramenter with the number of messages is required"
	}

	expectedMessages, ok := params[0].(int)
	if !ok || expectedMessages <= 0 {
		return "Number of messages should be a positive integer"
	}

	for i := 0; i < expectedMessages; i++ {
		res := p.shouldReceive(actual, nil, nil)
		if res != "" {
			return res
		}
	}

	return ""
}

func (p *Pact) ShouldStop(param1 interface{}, _ ...interface{}) string {
	actual, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	return p.shouldStop(actual)
}

// Should send one given message.
// Who is the receiver does not matter.
func (p *Pact) ShouldSend(param1 interface{}, params ...interface{}) string {
	actual, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	// If there is only one argument than it's the message to assert
	if len(params) != 1 {
		return "One or two paremeters are required to assert sending"
	}

	expectedMsg := params[0]

	return p.shouldSend(actual, nil, expectedMsg)
}

// Should send one given message to the specified receiver.
func (p *Pact) ShouldSendTo(param1 interface{}, params ...interface{}) string {
	actual, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	if len(params) != 2 {
		return "Too paremeters are required to assert sending"
	}

	expectedMsg := params[0]

	// If there are two arguments than the second is the expected target of sending
	receiver, ok := params[1].(*actor.PID)
	if !ok {
		return "Receiver should be an actor PID"
	}

	return p.shouldSend(actual, receiver, expectedMsg)
}

func (p *Pact) ShouldSendSomething(param1 interface{}, _ ...interface{}) string {
	actual, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	return p.shouldSend(actual, nil, nil)
}

// Should send N any messages
func (p *Pact) ShouldSendN(param1 interface{}, params ...interface{}) string {
	actual, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	if len(params) != 1 {
		return "One paramenter with the number of messages is required"
	}

	expectedMessages, ok := params[0].(int)
	if !ok || expectedMessages <= 0 {
		return "Number of messages should be a positive integer"
	}

	for i := 0; i < expectedMessages; i++ {
		res := p.shouldSend(actual, nil, nil)
		if res != "" {
			return res
		}
	}

	return ""
}

// TODO: Add a timeout parameter.
//       Otherwise this will not work for long running "reactions".
func (p *Pact) ShouldNotSendOrReceive(param1 interface{}, _ ...interface{}) string {
	actual, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	return p.shouldNotSendOrReceive(actual)
}
