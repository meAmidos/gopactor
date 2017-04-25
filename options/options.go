// Configuration options that are used by Pact
// when spawning actors. You can build a custom set of
// options for every actor you test with Pact.
//
// Exmaple
//
//   // The actor should use prefix "sender"
//   // and Pact should only intercept outbound messages
//   // for this actor.
//   // In addition, allow extra time when waiting for the actor to send something.
//   opt1 := OptOutboundInterceptionOnly.WithPrefix("sender").WithTimeout(time.Second)
//   actor1, _ := SpawnFromInstance(&MyActor{}, opt1)
//
//   // Pact should listen only to system messages received by the actor
//   opt2 := OptNoInterception.WithOutboundInterception()
//   actor2, _ := SpawnFromInstance(&MyActor{}, opt2)
//
//   // Use the default configuration:
//   // inbound/outbound interception and no interception of system messages.
//   // In addition, reduce the timeout (default value id 3 milliseconds).
//   opt3 := OptDefault.WithTimeout(1 * time.Millisecond)
//   actor3, _ := SpawnFromInstance(&MyActor{}, opt3)
package options

import "time"

// This timeout is used when waiting for inbound/outbound messages.
// Three milliseconds is long enough to allow for some regular operations
// in an actor to complete and for messages to be emitted. At the same
// time, it is short enough to keep testing reasonably fast even when
// you use hundreds of assertions.
const DEFAULT_TIMEOUT = 3 * time.Millisecond

// For each actor you test with Pact, you can define a custom set of
// options. The most useful ones are those that allow enabling
// interception selectively.
type Options struct {
	EnableInboundInterception  bool
	EnableOutboundInterception bool
	EnableSystemInterception   bool

	// A prefix of the spawned actor name.
	// It is useful mostly in cases when you debug your application
	// and want the actor's PID to have a meaningful value.
	Prefix string

	// The maximum amount of time to wait for an expected message.
	// This applies to all assertions for a given actor.
	Timeout time.Duration
}

// Predefined configuration: interception is disabled.
var OptNoInterception = Options{
	Timeout: DEFAULT_TIMEOUT,
}

// Predefined configuration for the most common scenario.
var OptDefault = Options{
	EnableInboundInterception:  true,
	EnableOutboundInterception: true,
	EnableSystemInterception:   false,
	Timeout:                    DEFAULT_TIMEOUT,
}

// Predefined configuration: intercept outbound messages only.
var OptOutboundInterceptionOnly = Options{
	EnableOutboundInterception: true,
	Timeout:                    DEFAULT_TIMEOUT,
}

// Predefined configuration: intercept inbound messages only.
var OptInboundInterceptionOnly = Options{
	EnableInboundInterception: true,
	Timeout:                   DEFAULT_TIMEOUT,
}

// A helper to add inbound interception to options
func (opt Options) WithInboundInterception() Options {
	opt.EnableInboundInterception = true
	return opt
}

// A helper to add outbound interception to options
func (opt Options) WithOutboundInterception() Options {
	opt.EnableOutboundInterception = true
	return opt
}

// A helper to add system messages interception to options
func (opt Options) WithSystemInterception() Options {
	opt.EnableSystemInterception = true
	return opt
}

// A helper to add prefix to options
func (opt Options) WithPrefix(prefix string) Options {
	opt.Prefix = prefix
	return opt
}

// A helper to add timeout to options
func (opt Options) WithTimeout(timeout time.Duration) Options {
	opt.Timeout = timeout
	return opt
}
