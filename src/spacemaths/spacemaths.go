package spacemaths

// GCD Greatest Common Denominator
func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple via GCD
func LCM(a, b uint64) uint64 {
	return a * b / GCD(a, b)
}
