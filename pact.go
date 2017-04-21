package pact

import (
	"fmt"
	"reflect"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

const (
	TEST_TIMEOUT = 3 * time.Millisecond
)

type Envelope struct {
	Sender  *actor.PID
	Target  *actor.PID
	Message interface{}
}

type Pact struct {
	SysCh         chan *Envelope
	InCh          chan *Envelope
	OutCh         chan *Envelope
	AssignedActor *actor.PID
	LoggingOn     bool
}

func New() *Pact {
	return &Pact{
		SysCh: make(chan *Envelope, 10),
		InCh:  make(chan *Envelope), // TODO: Use a buffered channel to not block on system messages
		OutCh: make(chan *Envelope),
	}
}

func (p *Pact) SpawnPrefix(obj actor.Actor, prefix string) (*actor.PID, error) {
	props := actor.
		FromInstance(obj).
		WithMiddleware(p.InboundMiddleware).
		WithOutboundMiddleware(p.OutboundMiddleware)

	// TODO: Spawn from a dedicated actor which could catch exceptions
	pid, err := actor.SpawnPrefix(props, prefix)
	if err != nil {
		return nil, err
	}

	p.AssignedActor = pid
	return pid, nil
}

func (p *Pact) shouldReceive(msg interface{}, from *actor.PID) string {
	select {
	case envelope := <-p.InCh:
		if msg == nil {
			return ""
		} else {
			return p.assertInboundMessage(envelope, msg, from)
		}
	case <-time.After(TEST_TIMEOUT):
		break
	}

	return "Timeout while waiting for a message"
}

func (p *Pact) shouldReceiveSysMsg(msg interface{}) string {
	for {
		select {
		case envelope := <-p.SysCh:
			if msg == nil { // Any message is ok
				return ""
			} else {
				// Ignore unmatching messages
				// This is important. Otherwise we would always have to check for
				// for the Start message first. And potentially other intermediate messages.
				match := p.assertInboundMessage(envelope, msg, nil)
				if match == "" {
					return ""
				}
			}
		case <-time.After(TEST_TIMEOUT):
			break
		}
	}

	return "Timeout while waiting for a system message"
}

func (p *Pact) shouldBeStopping() string {
	return p.shouldReceiveSysMsg(&actor.Stopping{})
}

func (p *Pact) assertInboundMessage(envelope *Envelope, msg interface{}, from *actor.PID) string {
	if !reflect.DeepEqual(envelope.Message, msg) {
		return fmt.Sprintf(`
Messages do not match
Expected: %#v
Actual: %#v
`, msg, envelope.Message)
	}

	if from != nil && !from.Equal(envelope.Sender) {
		return fmt.Sprintf(`
Sender does not match
Expected: %#v
Actual: %#v
`, from, envelope.Sender)
	}

	return ""
}

func (p *Pact) shouldSend(msg interface{}, target *actor.PID) string {
	select {
	case envelope := <-p.OutCh:
		if msg == nil {
			return ""
		} else {
			return p.assertOutboundMessage(envelope, msg, target)
		}
	case <-time.After(TEST_TIMEOUT):
		break
	}

	return "Timeout while waiting for sending"
}

func (p *Pact) assertOutboundMessage(envelope *Envelope, msg interface{}, target *actor.PID) string {
	if !reflect.DeepEqual(envelope.Message, msg) {
		return fmt.Sprintf(`
Messages do not match
Expected: %#v
Actual: %#v
`, msg, envelope.Message)
	}

	if target != nil && !target.Equal(envelope.Target) {
		return "Receiver does not match"
	}

	return ""
}

func (p *Pact) waitForAnything() string {
	select {
	case envelope := <-p.OutCh:
		return fmt.Sprintf("Got outbound message: %#v", envelope.Message)
	case envelope := <-p.InCh:
		return fmt.Sprintf("Got inbound message: %#v", envelope.Message)
	case <-time.After(TEST_TIMEOUT):
		break
	}

	return ""
}
