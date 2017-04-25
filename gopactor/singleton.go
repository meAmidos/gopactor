package gopactor

// DEFAULT_GOPACTOR is the default instance of Gopactor
// that is used when you invoke functions of the Gopactor package.
var DEFAULT_GOPACTOR *Gopactor

func init() {
	DEFAULT_GOPACTOR = New()
}
