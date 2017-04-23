package catcher

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
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

	// One followed actor per catcher
	AssignedActor *actor.PID

	LoggingOn bool
	options   Options
}

func (catcher *Catcher) id() string {
	if catcher.AssignedActor != nil {
		return catcher.AssignedActor.String()
	}

	return "-"
}

func New() *Catcher {
	return &Catcher{
		ChSystemInbound: make(chan *Envelope, 10),

		// These are deliberately not buffered to make synchronization points
		ChUserInbound:  make(chan *Envelope),
		ChUserOutbound: make(chan *Envelope),
	}
}

func (catcher *Catcher) Spawn(props *actor.Props, options ...Options) (*actor.PID, error) {
	var opt Options
	if len(options) == 0 {
		opt = OptDefault
	} else {
		opt = options[0]
	}

	catcher.options = opt

	if opt.EnableInboundInterception || opt.EnableSystemInterception {
		props = props.WithMiddleware(catcher.InboundMiddleware)
	}

	if opt.EnableOutboundInterception {
		props = props.WithOutboundMiddleware(catcher.OutboundMiddleware)
	}

	pid, err := actor.SpawnPrefix(props, opt.Prefix)
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
	case <-time.After(catcher.options.Timeout):
		return fmt.Sprintf("Timeout %s while waiting for a message", catcher.options.Timeout)
	}
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
		case <-time.After(catcher.options.Timeout):
			return fmt.Sprintf("Timeout %s while waiting for a system message", catcher.options.Timeout)
		}
	}
}

func (catcher *Catcher) ShouldSend(receiver *actor.PID, msg interface{}) string {
	select {
	case envelope := <-catcher.ChUserOutbound:
		if msg == nil { // Any message wil suffice
			return ""
		} else {
			return assertOutboundMessage(envelope, msg, receiver)
		}
	case <-time.After(catcher.options.Timeout):
		return fmt.Sprintf("Timeout %s while waiting for sending", catcher.options.Timeout)
	}
}

func (catcher *Catcher) ShouldNotSendOrReceive(pid *actor.PID) string {
	select {
	case envelope := <-catcher.ChUserOutbound:
		return fmt.Sprintf("Got outbound message: %#v", envelope.Message)
	case envelope := <-catcher.ChUserInbound:
		return fmt.Sprintf("Got inbound message: %#v", envelope.Message)
	case <-time.After(catcher.options.Timeout):
		return ""
	}
}
