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
