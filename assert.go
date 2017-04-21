package pact

import (
	"fmt"
	"reflect"

	"github.com/AsynkronIT/protoactor-go/actor"
)

func assertInboundMessage(envelope *Envelope, msg interface{}, sender *actor.PID) string {
	if !reflect.DeepEqual(envelope.Message, msg) {
		return fmt.Sprintf(`
Messages do not match
Expected: %#v
Actual: %#v
`, msg, envelope.Message)
	}

	if sender != nil && !sender.Equal(envelope.Sender) {
		return fmt.Sprintf(`
Sender does not match
Expected: %#v
Actual: %#v
`, sender, envelope.Sender)
	}

	return ""
}

func assertOutboundMessage(envelope *Envelope, msg interface{}, receiver *actor.PID) string {
	if !reflect.DeepEqual(envelope.Message, msg) {
		return fmt.Sprintf(`
Messages do not match
Expected: %#v
Actual: %#v
`, msg, envelope.Message)
	}

	if receiver != nil && !receiver.Equal(envelope.Target) {
		return "Receiver does not match"
	}

	return ""
}
