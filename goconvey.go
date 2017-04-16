package pact

import "github.com/AsynkronIT/protoactor-go/actor"

func (p *Pact) checkActualObject(actual interface{}) string {
	object, ok := actual.(*actor.PID)
	if !ok {
		return "Object is not an actor PID"
	}

	if !p.AssignedActor.Equal(object) {
		return "Object is not registered for tests"
	}

	return ""
}

func (p *Pact) ShouldReceive(actual interface{}, expected ...interface{}) string {
	p.checkActualObject(actual)

	// If a single argument is provided than it's a message
	//    and the sender does not matter
	if len(expected) == 1 {
		return p.shouldReceive(expected[0], nil)
	}

	if len(expected) != 2 {
		return "One or two paremeters are required to assert receiving"
	}

	// Two arguments means that the second is the expected sender
	from, ok := expected[1].(*actor.PID)
	if !ok {
		return "Sender should be an actor PID"
	}

	return p.shouldReceive(expected[0], from)
}

func (p *Pact) ShouldReceiveSomething(actual interface{}, _ ...interface{}) string {
	p.checkActualObject(actual)

	return p.shouldReceive(nil, nil)
}

func (p *Pact) ShouldSend(actual interface{}, expected ...interface{}) string {
	p.checkActualObject(actual)

	// If there is only one argument than it's the message to assert
	if len(expected) == 1 {
		return p.shouldSend(expected[0], nil)
	}

	if len(expected) != 2 {
		return "One or two paremeters are required to assert sending"
	}

	// If there are two arguments than the second is the expected target of sending
	target, ok := expected[1].(*actor.PID)
	if !ok {
		return "Receiver should be an actor PID"
	}

	return p.shouldSend(expected[0], target)
}

func (p *Pact) ShouldSendSomething(actual interface{}, _ ...interface{}) string {
	p.checkActualObject(actual)

	return p.shouldSend(nil, nil)
}

// TODO: Add a timeout parameter.
//       Otherwise this will not work for long running "reactions".
func (p *Pact) ShouldNotReact(actual interface{}, _ ...interface{}) string {
	p.checkActualObject(actual)

	return p.waitForAnything()
}
