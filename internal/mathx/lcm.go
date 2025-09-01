package mathx

// LCM returns the least common multiple of a and b
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	g := GCD(a, b)
	return abs(a*b) / g
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
