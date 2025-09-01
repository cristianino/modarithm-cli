package mathx

// Congruence checks if a ≡ b (mod m) and returns canonical residues
func Congruence(a, b, m int) (eq bool, ra, rb int) {
	if m == 0 {
		return a == b, a, b
	}
	ra = ((a % m) + m) % m
	rb = ((b % m) + m) % m
	eq = ra == rb
	return
}
