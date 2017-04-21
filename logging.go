package pact

import "fmt"

func (catcher *Catcher) StartLogging() {
	catcher.LoggingOn = true
}

func (catcher *Catcher) StopLogging() {
	catcher.LoggingOn = false
}

func (catcher *Catcher) TryLogMessage(text string, msg interface{}) {
	if catcher.LoggingOn {
		fmt.Printf(`
%s
=====================
 %#v
 ====================
`, text, msg)
	}

}
