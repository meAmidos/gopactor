package catcher

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/AsynkronIT/protoactor-go/actor"
)

func messagesMatch(actual, expected interface{}) bool {
	// Special case: compare Terminated messages
	// For the sake of simplicity, leave aside the AddressTerminated field of the messages.
	if termActual, ok := actual.(*actor.Terminated); ok {
		if termExpected, ok := expected.(*actor.Terminated); ok {
			// Any Terminated message will suffice
			if termExpected.Who == nil {
				return true
			}

			return termActual.Who.String() == termExpected.Who.String()
		}
	}

	return reflect.DeepEqual(actual, expected)
}

func assertInboundMessage(envelope *Envelope, msg interface{}, sender *actor.PID) string {
	if !messagesMatch(envelope.Message, msg) {
		return fmt.Sprintf(`
Messages do not match
Expected: %#v
Actual: %#v
`, msg, envelope.Message)
	}

	if sender != nil {
		if envelope.Sender == nil {
			return fmt.Sprintf(`
Sender is unknown
Expected: %#v
Actual: nil
`, sender)
		} else if !sender.Equal(envelope.Sender) {
			return fmt.Sprintf(`
Sender does not match
Expected: %#v
Actual: %#v
`, sender, envelope.Sender)
		}
	}

	return ""
}

func assertOutboundMessage(envelope *Envelope, msg interface{}, receiver *actor.PID) string {
	if !messagesMatch(envelope.Message, msg) {
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

func assertSpawnedActor(pid *actor.PID, match string) string {
	if !strings.Contains(pid.String(), match) {
		return fmt.Sprintf(`
The spawned actor's PID does not match
Expected: %s
Actual: %s
`, match, pid)
	}

	return ""
}
