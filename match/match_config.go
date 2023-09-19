package match

var maxRuneBytes = [...]byte{244, 143, 191, 191}

const (
	rightNoMatch result = iota
	rightMatch
	rightStop
)
