package pact

import "github.com/meamidos/pact/assertions"

var (
	ShouldReceive          = assertions.ShouldReceive
	ShouldReceiveFrom      = assertions.ShouldReceiveFrom
	ShouldReceiveSomething = assertions.ShouldReceiveSomething
	ShouldReceiveN         = assertions.ShouldReceiveN

	ShouldSend          = assertions.ShouldSend
	ShouldSendTo        = assertions.ShouldSendTo
	ShouldSendSomething = assertions.ShouldSendSomething
	ShouldSendN         = assertions.ShouldSendN

	ShouldNotSendOrReceive = assertions.ShouldNotSendOrReceive

	ShouldStart        = assertions.ShouldStart
	ShouldStop         = assertions.ShouldStop
	ShouldBeRestarting = assertions.ShouldBeRestarting
)
