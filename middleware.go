package pact

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

func (p *Pact) InboundMiddleware(next actor.ActorFunc) actor.ActorFunc {
	return func(ctx actor.Context) {
		p.ProcessInboundMessage(ctx)

		// Swap the context with it's thin wrapper which intercepts some calls.
		if _, ok := ctx.(*Context); !ok {
			ctx = NewContext(ctx)
		}
		next(ctx)
	}
}

func (p *Pact) ProcessInboundMessage(ctx actor.Context) {
	p.TryLogMessage("Received", ctx.Message())
	envelope := &Envelope{
		Sender:  ctx.Sender(),
		Target:  ctx.Self(),
		Message: ctx.Message(),
	}

	if !IsSystemMessage(ctx.Message()) {
		p.InCh <- envelope
	} else {
		p.SysCh <- envelope
	}
}

func (p *Pact) OutboundMiddleware(next actor.SenderFunc) actor.SenderFunc {
	return func(ctx actor.Context, target *actor.PID, env actor.MessageEnvelope) {
		p.ProcessOutboundMessage(ctx, target, env)
		next(ctx, target, env)
	}
}

func (p *Pact) ProcessOutboundMessage(ctx actor.Context, target *actor.PID, env actor.MessageEnvelope) {
	p.TryLogMessage("Sent", ctx.Message())
	if !IsSystemMessage(ctx.Message()) {
		p.OutCh <- &Envelope{
			Sender:  ctx.Self(),
			Target:  target,
			Message: env.Message,
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
