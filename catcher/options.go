package catcher

type Options struct {
	EnableInboundInterception  bool
	EnableOutboundInterception bool
	EnableSystemInterception   bool
}

var OptNoInterception = Options{}

var OptDefault = Options{
	EnableInboundInterception:  true,
	EnableOutboundInterception: true,
	EnableSystemInterception:   false,
}

var OptOutboundInterceptionOnly = Options{
	EnableOutboundInterception: true,
}

var OptInboundInterceptionOnly = Options{
	EnableInboundInterception: true,
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
