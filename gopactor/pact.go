package gopactor

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/meamidos/gopactor/catcher"
)

type Gopactor struct {
	CatchersByPID map[string]*catcher.Catcher
	LoggingOn     bool
}

func New() *Gopactor {
	p := &Gopactor{}
	p.Reset()
	return p
}

func (p *Gopactor) Reset() {
	p.CatchersByPID = make(map[string]*catcher.Catcher)
}

func (p *Gopactor) getCatcherByPID(pid *actor.PID) *catcher.Catcher {
	return p.CatchersByPID[pid.String()]
}

func (p *Gopactor) shouldReceive(receiver, sender *actor.PID, msg interface{}) string {
	catcher := p.getCatcherByPID(receiver)
	if catcher == nil {
		return "Receiver is not registered in Gopactor"
	}

	return catcher.ShouldReceive(sender, msg)
}

func (p *Gopactor) shouldReceiveSysMsg(receiver *actor.PID, msg interface{}) string {
	catcher := p.getCatcherByPID(receiver)
	if catcher == nil {
		return "Receiver is not registered in Gopactor"
	}

	return catcher.ShouldReceiveSysMsg(msg)
}

func (p *Gopactor) shouldStart(pid *actor.PID) string {
	return p.shouldReceiveSysMsg(pid, &actor.Started{})
}

func (p *Gopactor) shouldStop(pid *actor.PID) string {
	return p.shouldReceiveSysMsg(pid, &actor.Stopped{})
}

func (p *Gopactor) shouldBeRestarting(pid *actor.PID) string {
	return p.shouldReceiveSysMsg(pid, &actor.Restarting{})
}

func (p *Gopactor) shouldSend(sender, receiver *actor.PID, msg interface{}) string {
	catcher := p.getCatcherByPID(sender)
	if catcher == nil {
		return "Sender is not registered in Gopactor"
	}

	return catcher.ShouldSend(receiver, msg)
}

func (p *Gopactor) shouldNotSendOrReceive(pid *actor.PID) string {
	catcher := p.getCatcherByPID(pid)
	if catcher == nil {
		return "Sender is not registered in Gopactor"
	}

	return catcher.ShouldNotSendOrReceive(pid)
}
