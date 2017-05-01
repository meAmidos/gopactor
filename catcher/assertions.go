package catcher

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/AsynkronIT/protoactor-go/actor"
)

func messagesMatch(msg1, msg2 interface{}) bool {
	return reflect.DeepEqual(msg1, msg2)
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
