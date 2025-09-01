package mathx

import "testing"

func TestMultiplesSym(t *testing.T) {
	negs, zero, poss := MultiplesSym(5, 2)
	wantNegs := []int{-10, -5}
	wantZero := 0
	wantPoss := []int{5, 10}
	if !equalSlice(negs, wantNegs) || zero != wantZero || !equalSlice(poss, wantPoss) {
		t.Errorf("MultiplesSym(5,2) = %v,%d,%v; want %v,%d,%v", negs, zero, poss, wantNegs, wantZero, wantPoss)
	}
}

func TestMultiplesAsym(t *testing.T) {
	negs, zero, poss := MultiplesAsym(3, 1, 2)
	wantNegs := []int{-3}
	wantZero := 0
	wantPoss := []int{3, 6}
	if !equalSlice(negs, wantNegs) || zero != wantZero || !equalSlice(poss, wantPoss) {
		t.Errorf("MultiplesAsym(3,1,2) = %v,%d,%v; want %v,%d,%v", negs, zero, poss, wantNegs, wantZero, wantPoss)
	}
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
