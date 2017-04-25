package gopactor

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// ShouldReceive is an assertion method. Its rules are:
// - The receiver should receive a given message.
// - It does not matter who is the sender.
func (p *Gopactor) ShouldReceive(param1 interface{}, params ...interface{}) string {
	receiver, ok := param1.(*actor.PID)
	if !ok {
		return "Receiver is not an actor PID"
	}

	if len(params) != 1 {
		return "One parameter with a message is required to assert receiving"
	}

	expectedMsg := params[0]

	return p.shouldReceive(receiver, nil, expectedMsg)
}

// ShouldReceiveFrom is an assertion method. Its rules are:
// - The receiver should receive a given message from a given sender.
func (p *Gopactor) ShouldReceiveFrom(param1 interface{}, params ...interface{}) string {
	receiver, ok := param1.(*actor.PID)
	if !ok {
		return "Receiver is not an actor PID"
	}

	if len(params) != 2 {
		return "Two parameters are required to assert receiving"
	}

	// Two arguments means that the second is the expected sender
	sender, ok := params[0].(*actor.PID)
	if !ok {
		return "Sender should be an actor PID"
	}

	expectedMsg := params[1]

	return p.shouldReceive(receiver, sender, expectedMsg)
}

// ShouldReceiveSomething is an assertion method. Its rules are:
// - The receiver should receive at least one message of any kind.
// - It does not matter who is the sender.
func (p *Gopactor) ShouldReceiveSomething(param1 interface{}, _ ...interface{}) string {
	receiver, ok := param1.(*actor.PID)
	if !ok {
		return "Receiver is not an actor PID"
	}

	return p.shouldReceive(receiver, nil, nil)
}

// ShouldReceiveN is an assertion method. Its rules are:
// - The receiver should receive N messages of any kind.
// - It does not matter who is the sender.
func (p *Gopactor) ShouldReceiveN(param1 interface{}, params ...interface{}) string {
	receiver, ok := param1.(*actor.PID)
	if !ok {
		return "Receiver is not an actor PID"
	}

	if len(params) != 1 {
		return "One paramenter with the number of expected messages is required"
	}

	expectedMessages, ok := params[0].(int)
	if !ok || expectedMessages <= 0 {
		return "Number of expected messages should be a positive integer"
	}

	for i := 0; i < expectedMessages; i++ {
		res := p.shouldReceive(receiver, nil, nil)
		if res != "" {
			return fmt.Sprintf("Expected %d messages, but got %d", expectedMessages, i)
		}
	}

	return ""
}

// ShouldStart is an assertion method. Its rules are:
// - The actor should receive a system message that indicates the actor has been started.
func (p *Gopactor) ShouldStart(param1 interface{}, _ ...interface{}) string {
	pid, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	return p.shouldStart(pid)
}

// ShouldStop is an assertion method. Its rules are:
// - The actor should receive a system message that indicates the actor has been stopped.
func (p *Gopactor) ShouldStop(param1 interface{}, _ ...interface{}) string {
	pid, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	return p.shouldStop(pid)
}

// ShouldBeRestarting is an assertion method. Its rules are:
// - The actor should receive a system message that indicates the actor is being restarted.
func (p *Gopactor) ShouldBeRestarting(param1 interface{}, _ ...interface{}) string {
	pid, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	return p.shouldBeRestarting(pid)
}

// ShouldSend is an assertion method. Its rules are:
// - The sender should send one given message.
// - It does not matter who is the receiver of the message.
func (p *Gopactor) ShouldSend(param1 interface{}, params ...interface{}) string {
	sender, ok := param1.(*actor.PID)
	if !ok {
		return "Sender is not an actor PID"
	}

	// If there is only one argument than it's the message to assert
	if len(params) != 1 {
		return "One parameter with a message is required to assert sending"
	}

	expectedMsg := params[0]

	return p.shouldSend(sender, nil, expectedMsg)
}

// ShouldSendTo is an assertion method. Its rules are:
// - The sender should send one given message to the specified receiver.
func (p *Gopactor) ShouldSendTo(param1 interface{}, params ...interface{}) string {
	sender, ok := param1.(*actor.PID)
	if !ok {
		return "Sender is not an actor PID"
	}

	if len(params) != 2 {
		return "Two parameters are required to assert sending"
	}

	// If there are two arguments than the second is the expected target of sending
	receiver, ok := params[0].(*actor.PID)
	if !ok {
		return "Receiver should be an actor PID"
	}

	expectedMsg := params[1]

	return p.shouldSend(sender, receiver, expectedMsg)
}

// ShouldSendSomething is an assertion method. Its rules are:
// - The sender should send at least one message of any kind.
// - It does not matter who is the receiver of the message.
func (p *Gopactor) ShouldSendSomething(param1 interface{}, _ ...interface{}) string {
	sender, ok := param1.(*actor.PID)
	if !ok {
		return "Sender is not an actor PID"
	}

	return p.shouldSend(sender, nil, nil)
}

// ShouldSendN is an assertion method. Its rules are:
// - The sender should send N messages of any kind.
// - It does not matter who are the receivers.
func (p *Gopactor) ShouldSendN(param1 interface{}, params ...interface{}) string {
	sender, ok := param1.(*actor.PID)
	if !ok {
		return "Sender is not an actor PID"
	}

	if len(params) != 1 {
		return "One paramenter with the number of expected messages is required"
	}

	expectedMessages, ok := params[0].(int)
	if !ok || expectedMessages <= 0 {
		return "Number of expected messages should be a positive integer"
	}

	for i := 0; i < expectedMessages; i++ {
		res := p.shouldSend(sender, nil, nil)
		if res != "" {
			return fmt.Sprintf("Expected %d messages to be sent, but got %d", expectedMessages, i)
		}
	}

	return ""
}

// ShouldNotSendOrReceive is an assertion method. Its rules are:
// - The actor should not send or receive any message.
func (p *Gopactor) ShouldNotSendOrReceive(param1 interface{}, _ ...interface{}) string {
	object, ok := param1.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	return p.shouldNotSendOrReceive(object)
}
