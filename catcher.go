package pact

import (
	"fmt"
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

type Catcher struct {
	ChSystemInbound chan *Envelope
	ChUserInbound   chan *Envelope
	ChUserOutbound  chan *Envelope

	// Channels dedicated to specific messages
	ChStopped chan *Envelope
	ChStarted chan *Envelope

	// One followed actor per catcher
	AssignedActor *actor.PID

	LoggingOn bool
}

func NewCatcher() *Catcher {
	return &Catcher{
		ChSystemInbound: make(chan *Envelope, 10),
		ChStopped:       make(chan *Envelope, 1),
		ChStarted:       make(chan *Envelope, 1),

		// These are deliberately not buffered to make synchronization points
		ChUserInbound:  make(chan *Envelope),
		ChUserOutbound: make(chan *Envelope),
	}
}

func (catcher *Catcher) Spawn(props *actor.Props, prefix string) (*actor.PID, error) {
	props = props.
		WithMiddleware(catcher.InboundMiddleware).
		WithOutboundMiddleware(catcher.OutboundMiddleware)

	pid, err := actor.SpawnPrefix(props, prefix)
	if err != nil {
		return nil, err
	}

	catcher.AssignedActor = pid
	return pid, nil
}

func (catcher *Catcher) ShouldReceive(sender *actor.PID, msg interface{}) string {
	select {
	case envelope := <-catcher.ChUserInbound:
		if msg == nil { // Any massage will suffice
			return ""
		} else {
			return assertInboundMessage(envelope, msg, sender)
		}
	case <-time.After(TEST_TIMEOUT):
		break
	}

	return "Timeout while waiting for a message"
}

func (catcher *Catcher) ShouldReceiveSysMsg(msg interface{}) string {
	for {
		select {
		case envelope := <-catcher.ChSystemInbound:
			if msg == nil { // Any message is ok
				return ""
			} else {
				// Ignore unmatching messages
				// This is important. Otherwise we would always have to check for
				// for the Start message first. And potentially for other intermediate messages.
				match := assertInboundMessage(envelope, msg, nil)
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

func (catcher *Catcher) ShouldSend(receiver *actor.PID, msg interface{}) string {
	select {
	case envelope := <-catcher.ChUserOutbound:
		if msg == nil { // Any message wil suffice
			return ""
		} else {
			return assertOutboundMessage(envelope, msg, receiver)
		}
	case <-time.After(TEST_TIMEOUT):
		break
	}

	return "Timeout while waiting for sending"
}

func (catcher *Catcher) ShouldNotSendOrReceive(pid *actor.PID) string {
	select {
	case envelope := <-catcher.ChUserOutbound:
		return fmt.Sprintf("Got outbound message: %#v", envelope.Message)
	case envelope := <-catcher.ChUserInbound:
		return fmt.Sprintf("Got inbound message: %#v", envelope.Message)
	case <-time.After(TEST_TIMEOUT):
		break
	}

	return ""
}
