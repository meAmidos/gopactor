package assertions

import "github.com/meamidos/pact/pact"

func ShouldReceive(actual interface{}, expected ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldReceive(actual, expected...)
}

func ShouldReceiveFrom(actual interface{}, expected ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldReceiveFrom(actual, expected...)
}

func ShouldReceiveSomething(actual interface{}, expected ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldReceiveSomething(actual, expected...)
}

func ShouldReceiveN(actual interface{}, params ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldReceiveN(actual, params...)
}

func ShouldSend(actual interface{}, expected ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldSend(actual, expected...)
}

func ShouldSendTo(actual interface{}, expected ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldSendTo(actual, expected...)
}

func ShouldSendSomething(actual interface{}, _ ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldSendSomething(actual)
}

func ShouldSendN(actual interface{}, params ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldSendN(actual, params...)
}

func ShouldNotSendOrReceive(actual interface{}, _ ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldNotSendOrReceive(actual)
}

func ShouldStart(actual interface{}, _ ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldStart(actual)
}

func ShouldStop(actual interface{}, _ ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldStop(actual)
}

func ShouldBeRestarting(actual interface{}, _ ...interface{}) string {
	return pact.DEFAULT_PACT.ShouldBeRestarting(actual)
}
