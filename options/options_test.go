package options_test

import (
	"testing"
	"time"

	"github.com/meamidos/gopactor/options"
	"github.com/stretchr/testify/assert"
)

func TestOptionsWith(t *testing.T) {
	a := assert.New(t)

	emptyOptions := options.Options{}

	// Empty options
	options := emptyOptions
	a.False(options.InboundInterceptionEnabled)
	a.False(options.OutboundInterceptionEnabled)
	a.False(options.SystemInterceptionEnabled)
	a.Empty(options.Prefix)
	a.Equal(time.Duration(0), options.Timeout)

	// With Prefix
	options = emptyOptions.WithPrefix("pref")
	a.False(options.InboundInterceptionEnabled)
	a.False(options.OutboundInterceptionEnabled)
	a.False(options.SystemInterceptionEnabled)
	a.Equal("pref", options.Prefix)
	a.Equal(time.Duration(0), options.Timeout)

	// With timeout
	options = emptyOptions.WithTimeout(time.Microsecond)
	a.False(options.InboundInterceptionEnabled)
	a.False(options.OutboundInterceptionEnabled)
	a.False(options.SystemInterceptionEnabled)
	a.Empty(options.Prefix)
	a.Equal(time.Microsecond, options.Timeout)

	// With inbound interception
	options = emptyOptions.WithInboundInterception()
	a.True(options.InboundInterceptionEnabled)
	a.False(options.OutboundInterceptionEnabled)
	a.False(options.SystemInterceptionEnabled)
	a.Empty(options.Prefix)
	a.Equal(time.Duration(0), options.Timeout)

	// With outbound interception
	options = emptyOptions.WithOutboundInterception()
	a.False(options.InboundInterceptionEnabled)
	a.True(options.OutboundInterceptionEnabled)
	a.False(options.SystemInterceptionEnabled)
	a.Empty(options.Prefix)
	a.Equal(time.Duration(0), options.Timeout)

	// With sys messages interception
	options = emptyOptions.WithSystemInterception()
	a.False(options.InboundInterceptionEnabled)
	a.False(options.OutboundInterceptionEnabled)
	a.True(options.SystemInterceptionEnabled)
	a.Empty(options.Prefix)
	a.Equal(time.Duration(0), options.Timeout)
}
