// Package options defines configuration options that are used by Gopactor
// when spawning actors. You can build a custom set of
// options for every actor you test with Gopactor.
//
// Example:
//
//   // Let's build a set of options with these requirements:
//   // - The actor should be spawned with prefix "sender".
//   // - Gopactor should only intercept outbound messages for this actor.
//   // - In addition, allow extra time when waiting for the actor to send something.
//   opt1 := OptOutboundInterceptionOnly.WithPrefix("sender").WithTimeout(time.Second)
//   actor1, _ := SpawnFromInstance(&MyActor{}, opt1)
//
//   // Another simple configuration:
//   // - Ask Gopactor to listen only to system messages received by the actor.
//   opt2 := OptNoInterception.WithOutboundInterception()
//   actor2, _ := SpawnFromInstance(&MyActor{}, opt2)
//
//   // Use the default configuration:
//   // - Inbound and outbound interception.
//   // - No interception of system messages.
//   // - In addition, reduce the timeout (default value id 3 milliseconds).
//   opt3 := OptDefault.WithTimeout(1 * time.Millisecond)
//   actor3, _ := SpawnFromInstance(&MyActor{}, opt3)
package options

import "time"

// DEFAULT_TIMEOUT value is used when no custom timeout has been specified.
// Three milliseconds is long enough to allow for some regular operations
// in an actor to complete and for messages to be emitted. At the same
// time, it is short enough to keep testing reasonably fast even when
// you use hundreds of assertions.
const DEFAULT_TIMEOUT = 3 * time.Millisecond

// Options is a container for individual configuration options used by Gopactor.
// For each actor you test with Gopactor, you can build a custom set of
// options.
type Options struct {
	// Which messages to intercept
	InboundInterceptionEnabled  bool
	OutboundInterceptionEnabled bool
	SystemInterceptionEnabled   bool

	// Spawning
	SpawnInterceptionEnabled bool
	DummySpawningEnabled     bool

	// A prefix of the spawned actor name.
	// It is useful mostly in cases when you debug your application
	// and want the actor's PID to have a meaningful value.
	Prefix string

	// The maximum amount of time to wait for an expected inbound or outbound message.
	// This applies to all assertions for a given actor.
	Timeout time.Duration
}

// OptNoInterception is one of predefined configurations:
// - interception is disabled
// - no dummy spawning
var OptNoInterception = Options{
	Timeout: DEFAULT_TIMEOUT,
}

// OptDefault is one of predefined configurations.
// It is used by Gopactor when no configuration is provided.
var OptDefault = Options{
	InboundInterceptionEnabled:  true,
	OutboundInterceptionEnabled: true,
	SystemInterceptionEnabled:   false,
	DummySpawningEnabled:        true,
	Timeout:                     DEFAULT_TIMEOUT,
}

// OptOutboundInterceptionOnly is one of predefined configurations:
// - intercept outbound messages only
// - dummy spawning is enabled
var OptOutboundInterceptionOnly = Options{
	OutboundInterceptionEnabled: true,
	DummySpawningEnabled:        true,
	Timeout:                     DEFAULT_TIMEOUT,
}

// OptInboundInterceptionOnly is one of predefined configurations:
// - intercept inbound messages only
// - dummy spawning is enabled
var OptInboundInterceptionOnly = Options{
	InboundInterceptionEnabled: true,
	DummySpawningEnabled:       true,
	Timeout:                    DEFAULT_TIMEOUT,
}

// WithInboundInterception is a helper method to add inbound interception to options
func (opt Options) WithInboundInterception() Options {
	opt.InboundInterceptionEnabled = true
	return opt
}

// WithOutboundInterception is a helper method to add outbound interception to options
func (opt Options) WithOutboundInterception() Options {
	opt.OutboundInterceptionEnabled = true
	return opt
}

// WithSystemInterception is a helper method to add system messages interception to options
func (opt Options) WithSystemInterception() Options {
	opt.SystemInterceptionEnabled = true
	return opt
}

// WithSpawnInterception is a helper method to add spawning interception to options
func (opt Options) WithSpawnInterception() Options {
	opt.SpawnInterceptionEnabled = true
	return opt
}

// WithDummySpawning is a helper method to add dummy spawning to options
func (opt Options) WithDummySpawning() Options {
	opt.DummySpawningEnabled = true
	return opt
}

// WithRealSpawning is a helper method to disable dummy spawning in options
func (opt Options) WithRealSpawning() Options {
	opt.DummySpawningEnabled = false
	return opt
}

// WithPrefix is a helper method to add prefix to options
func (opt Options) WithPrefix(prefix string) Options {
	opt.Prefix = prefix
	return opt
}

// WithTimeout is a helper to add timeout to options
func (opt Options) WithTimeout(timeout time.Duration) Options {
	opt.Timeout = timeout
	return opt
}
