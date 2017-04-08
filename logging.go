package pact

import "fmt"

func (p *Pact) StartLogging() {
	p.LoggingOn = true
}

func (p *Pact) StopLogging() {
	p.LoggingOn = false
}

func (p *Pact) TryLogMessage(text string, msg interface{}) {
	if p.LoggingOn {
		fmt.Printf(`
%s
=====================
 %#v
 ====================
`, text, msg)
	}

}
