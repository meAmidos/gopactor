package catcher

import "time"

const DEFAULT_TIMEOUT = 3 * time.Millisecond

type Options struct {
	EnableInboundInterception  bool
	EnableOutboundInterception bool
	EnableSystemInterception   bool
	Prefix                     string
	Timeout                    time.Duration
}

var OptNoInterception = Options{
	Timeout: DEFAULT_TIMEOUT,
}

var OptDefault = Options{
	EnableInboundInterception:  true,
	EnableOutboundInterception: true,
	EnableSystemInterception:   false,
	Timeout:                    DEFAULT_TIMEOUT,
}

var OptOutboundInterceptionOnly = Options{
	EnableOutboundInterception: true,
	Timeout:                    DEFAULT_TIMEOUT,
}

var OptInboundInterceptionOnly = Options{
	EnableInboundInterception: true,
	Timeout:                   DEFAULT_TIMEOUT,
}

func (opt Options) WithInboundInterception() Options {
	opt.EnableInboundInterception = true
	return opt
}

func (opt Options) WithOutboundInterception() Options {
	opt.EnableOutboundInterception = true
	return opt
}

func (opt Options) WithSystemInterception() Options {
	opt.EnableSystemInterception = true
	return opt
}

func (opt Options) WithPrefix(prefix string) Options {
	opt.Prefix = prefix
	return opt
}

func (opt Options) WithTimeout(timeout time.Duration) Options {
	opt.Timeout = timeout
	return opt
}
