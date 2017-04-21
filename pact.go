package pact

import (
	"github.com/AsynkronIT/protoactor-go/actor"
)

type Pact struct {
	CatchersByPID map[string]*Catcher
	LoggingOn     bool
}

func New() *Pact {
	return &Pact{
		CatchersByPID: make(map[string]*Catcher),
	}
}

var DEFAULT_PACT *Pact

func init() {
	DEFAULT_PACT = New()
}

func (p *Pact) GetCatcherByPID(pid *actor.PID) *Catcher {
	return p.CatchersByPID[pid.String()]
}

func (p *Pact) SpawnFromInstance(obj actor.Actor, prefix string) (*actor.PID, error) {
	catcher := NewCatcher()
	catcher.LoggingOn = p.LoggingOn

	pid, err := catcher.SpawnFromInstance(obj, prefix)
	if err != nil {
		return nil, err
	}

	p.CatchersByPID[pid.String()] = catcher

	return pid, nil
}

func (p *Pact) shouldReceive(receiver, sender *actor.PID, msg interface{}) string {
	catcher := p.GetCatcherByPID(receiver)
	if catcher == nil {
		return "Receiver is not registered in Pact"
	}

	return catcher.ShouldReceive(sender, msg)
}

func (p *Pact) shouldReceiveSysMsg(receiver *actor.PID, msg interface{}) string {
	catcher := p.GetCatcherByPID(receiver)
	if catcher == nil {
		return "Receiver is not registered in Pact"
	}

	return catcher.ShouldReceiveSysMsg(msg)
}

func (p *Pact) shouldStop(pid *actor.PID) string {
	return p.shouldReceiveSysMsg(pid, &actor.Stopped{})
}

func (p *Pact) shouldSend(sender, receiver *actor.PID, msg interface{}) string {
	catcher := p.GetCatcherByPID(receiver)
	if catcher == nil {
		return "Sender is not registered in Pact"
	}

	return catcher.ShouldSend(receiver, msg)
}

func (p *Pact) shouldNotSendOrReceive(pid *actor.PID) string {
	catcher := p.GetCatcherByPID(pid)
	if catcher == nil {
		return "Sender is not registered in Pact"
	}

	return catcher.ShouldNotSendOrReceive(pid)
}
