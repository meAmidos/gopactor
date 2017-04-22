package catcher

type Options struct {
	EnableInboundInterception  bool
	EnableOutboundInterception bool
}

var OptNoInterception = Options{}

var OptDefault = Options{
	EnableInboundInterception:  true,
	EnableOutboundInterception: true,
}

var OptOutboundInterceptionOnly = Options{
	EnableOutboundInterception: true,
}

var OptInboundInterceptionOnly = Options{
	EnableInboundInterception: true,
}
