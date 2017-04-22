package pact

var DEFAULT_PACT *Pact

func init() {
	DEFAULT_PACT = New()
}
