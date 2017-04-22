package catcher

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

func (catcher *Catcher) InboundMiddleware(next actor.ActorFunc) actor.ActorFunc {
	return func(ctx actor.Context) {
		catcher.ProcessInboundMessage(ctx)

		// Swap the context with it's thin wrapper which intercepts some calls.
		if _, ok := ctx.(*Context); !ok {
			ctx = NewContext(ctx)
		}
		next(ctx)
	}
}

func (catcher *Catcher) ProcessInboundMessage(ctx actor.Context) {
	message := ctx.Message()

	catcher.TryLogMessage("Received", message)
	envelope := &Envelope{
		Sender:  ctx.Sender(),
		Target:  ctx.Self(),
		Message: message,
	}

	if !IsSystemMessage(message) {
		// fmt.Printf("\n== INBOUND (%s): %#v\n", catcher.id(), message)
		catcher.ChUserInbound <- envelope
	} else {
		// fmt.Printf("\n== IN SYS (%s): %#v\n", catcher.id(), message)
		catcher.ProcessSystemMessage(envelope)
	}
}

func (catcher *Catcher) ProcessSystemMessage(envelope *Envelope) {
	catcher.ChSystemInbound <- envelope
}

func (catcher *Catcher) OutboundMiddleware(next actor.SenderFunc) actor.SenderFunc {
	return func(ctx actor.Context, target *actor.PID, env actor.MessageEnvelope) {
		catcher.ProcessOutboundMessage(ctx, target, env)
		next(ctx, target, env)
	}
}

// TODO: Is there a difference between using ctx.Message() and env.Message?
func (catcher *Catcher) ProcessOutboundMessage(ctx actor.Context, target *actor.PID, env actor.MessageEnvelope) {
	message := env.Message

	catcher.TryLogMessage("Sent", message)
	if !IsSystemMessage(message) {
		catcher.ChUserOutbound <- &Envelope{
			Sender:  ctx.Self(),
			Target:  target,
			Message: message,
		}
	}
}

func IsSystemMessage(msg interface{}) bool {
	switch msg.(type) {
	case actor.AutoReceiveMessage:
		return true
	case actor.SystemMessage:
		return true
	}

	return false
}
