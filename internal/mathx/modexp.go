package mathx

// ModExp computes base^exp mod m using exponentiation by squaring
func ModExp(base, exp, m int) int {
	if m == 0 {
		return 0
	}
	result := 1
	base = base % m
	if base < 0 {
		base += m
	}
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % m
		}
		exp /= 2
		base = (base * base) % m
	}
	return result
}
