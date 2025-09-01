package mathx

import "testing"

func TestCongruence(t *testing.T) {
	cases := []struct {
		a, b, m int
		eq      bool
		ra, rb  int
	}{
		{45, 9, 6, true, 3, 3},
		{10, 3, 7, true, 3, 3},
		{14, 14, 5, true, 4, 4},
		{0, 0, 1, true, 0, 0},
		{7, 2, 0, false, 7, 2},
	}
	for _, c := range cases {
		eq, ra, rb := Congruence(c.a, c.b, c.m)
		if eq != c.eq || ra != c.ra || rb != c.rb {
			t.Errorf("Congruence(%d,%d,%d)=%v,%d,%d; want %v,%d,%d", c.a, c.b, c.m, eq, ra, rb, c.eq, c.ra, c.rb)
		}
	}
}
