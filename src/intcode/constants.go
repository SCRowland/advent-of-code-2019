package intcode

const (
	add           = 1
	multiply      = 2
	input         = 3
	output        = 4
	jumpIfTrue    = 5
	jumpIfFalse   = 6
	lessThan      = 7
	equals        = 8
	adjustRelBase = 9
	halt          = 99
)

const (
	positional = 0
	immediate  = 1
	relative   = 2
)
