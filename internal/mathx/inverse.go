package mathx

// ModInverse returns the modular inverse of a mod m, if it exists
func ModInverse(a, m int) (inv int, ok bool) {
	g, x, _ := extendedGCD(a, m)
	if g != 1 {
		return 0, false
	}
	inv = x % m
	if inv < 0 {
		inv += m
	}
	return inv, true
}

func extendedGCD(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := extendedGCD(b, a%b)
	x, y = y1, x1-(a/b)*y1
	return
}
