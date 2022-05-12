package parameters

type AbortSource uint

const (
	DialogServiceUser AbortSource = iota
	DialogServiceProvider
)
