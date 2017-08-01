package catcher

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

func (catcher *Catcher) inboundMiddleware(next actor.ActorFunc) actor.ActorFunc {
	return func(ctx actor.Context) {
		catcher.processInboundMessage(ctx)

		// Swap the context with a thin wrapper which intercepts some calls.
		if _, ok := ctx.(*Context); !ok {
			ctx = NewContext(catcher, ctx)
		}
		next(ctx)
	}
}

func (catcher *Catcher) processInboundMessage(ctx actor.Context) {
	message := ctx.Message()

	envelope := &Envelope{
		Sender:  ctx.Sender(),
		Target:  ctx.Self(),
		Message: message,
	}

	if !isSystemMessage(message) {
		if catcher.Options.InboundInterceptionEnabled {
			catcher.ChUserInbound <- envelope
		}
	} else {
		if catcher.Options.SystemInterceptionEnabled {
			catcher.processSystemMessage(envelope)
		}
	}
}

func (catcher *Catcher) processSystemMessage(envelope *Envelope) {
	catcher.ChSystemInbound <- envelope
}

func (catcher *Catcher) outboundMiddleware(next actor.SenderFunc) actor.SenderFunc {
	fn := actor.SenderFunc(func(ctx actor.Context, target *actor.PID, env actor.MessageEnvelope) {
		catcher.processOutboundMessage(ctx, target, env)
		next(ctx, target, &env)
	})

	return fn
	/*	return func(ctx actor.Context, target *actor.PID, env actor.MessageEnvelope)  {
		catcher.processOutboundMessage(ctx, target, env)
		next(ctx, target, env)
	}*/
}

func (catcher *Catcher) processOutboundMessage(ctx actor.Context, target *actor.PID, env actor.MessageEnvelope) {
	// TODO: Is there a difference between using ctx.Message() and env.Message?
	message := env.Message

	if !isSystemMessage(message) {
		catcher.ChUserOutbound <- &Envelope{
			Sender:  ctx.Self(),
			Target:  target,
			Message: message,
		}
	}
}

func isSystemMessage(msg interface{}) bool {
	switch msg.(type) {
	case actor.AutoReceiveMessage:
		return true
	case actor.SystemMessage:
		return true
	}

	return false
}
