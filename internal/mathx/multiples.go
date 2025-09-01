package mathx

// MultiplesSym returns negatives, zero, positives for symmetric limit
func MultiplesSym(n, limit int) (negs []int, zero int, poss []int) {
	zero = 0
	for i := -limit; i < 0; i++ {
		negs = append(negs, i*n)
	}
	for i := 1; i <= limit; i++ {
		poss = append(poss, i*n)
	}
	return
}

// MultiplesAsym returns negatives, zero, positives for asymmetric limits
func MultiplesAsym(n, limitNeg, limitPos int) (negs []int, zero int, poss []int) {
	zero = 0
	for i := -limitNeg; i < 0; i++ {
		negs = append(negs, i*n)
	}
	for i := 1; i <= limitPos; i++ {
		poss = append(poss, i*n)
	}
	return
}
