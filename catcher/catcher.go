package catcher

import (
	"fmt"
	"strings"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"

	"gopactor/options"
)

// Envelope is a wrapper that is used for all intercepted messages
type Envelope struct {
	Sender  *actor.PID
	Target  *actor.PID
	Message interface{}
}

// Catcher is the working horse of the interception mechanism.
// It seats in front of every tested actor and watches for
// messages and system events.
type Catcher struct {
	// Channels for intercepted messages
	ChSystemInbound chan *Envelope
	ChUserInbound   chan *Envelope
	ChUserOutbound  chan *Envelope

	// Channels for intercepted spawning of children
	ChSpawning chan *actor.PID

	// One followed actor per catcher
	AssignedActor *actor.PID

	Options options.Options
}

// This is used for logging purposes only
func (catcher *Catcher) id() string {
	if catcher.AssignedActor != nil {
		return catcher.AssignedActor.String()
	}

	return "-"
}

// New creates a new instance of Catcher.
func New() *Catcher {
	return &Catcher{
		ChSystemInbound: make(chan *Envelope, 10),

		// These are deliberately not buffered to make synchronization points
		ChUserInbound:  make(chan *Envelope),
		ChUserOutbound: make(chan *Envelope),
		ChSpawning:     make(chan *actor.PID),
	}
}

// Spawn an actor with injected middleware.
func (catcher *Catcher) Spawn(props *actor.Props, opts ...options.Options) (*actor.PID, error) {
	var opt options.Options
	if len(opts) == 0 {
		opt = options.OptDefault
	} else {
		opt = opts[0]
	}

	catcher.Options = opt

	if opt.InboundInterceptionEnabled || opt.SystemInterceptionEnabled || opt.SpawnInterceptionEnabled || opt.DummySpawningEnabled {
		props = props.WithMiddleware(catcher.inboundMiddleware)
	}

	if opt.OutboundInterceptionEnabled {
		props = props.WithOutboundMiddleware(catcher.outboundMiddleware)
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
	case <-time.After(catcher.Options.Timeout):
		return fmt.Sprintf("Timeout %s while waiting for a message", catcher.Options.Timeout)
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
		case <-time.After(catcher.Options.Timeout):
			return fmt.Sprintf("Timeout %s while waiting for a system message", catcher.Options.Timeout)
		}
	}
}

func (catcher *Catcher) ShouldSend(receiver *actor.PID, msg interface{}) string {
	select {
	case envelope := <-catcher.ChUserOutbound:
		if msg == nil { // Any message will suffice
			return ""
		} else {
			return assertOutboundMessage(envelope, msg, receiver)
		}
	case <-time.After(catcher.Options.Timeout):
		return fmt.Sprintf("Timeout %s while waiting for sending", catcher.Options.Timeout)
	}
}

func (catcher *Catcher) ShouldNotSendOrReceive(pid *actor.PID) string {
	select {
	case envelope := <-catcher.ChUserOutbound:
		return fmt.Sprintf("Got outbound message: %#v", envelope.Message)
	case envelope := <-catcher.ChUserInbound:
		return fmt.Sprintf("Got inbound message: %#v", envelope.Message)
	case <-time.After(catcher.Options.Timeout):
		return ""
	}
}

func (catcher *Catcher) ShouldSpawn(match string) string {
	select {
	case pid := <-catcher.ChSpawning:
		if match == "" { // Any spawned actor will suffice
			return ""
		} else {
			if strings.Contains(pid.String(), match) {
				return ""
			} else {
				return assertSpawnedActor(pid, match)
			}
		}
	case <-time.After(catcher.Options.Timeout):
		return fmt.Sprintf("Timeout %s while waiting for spawning", catcher.Options.Timeout)
	}
}
