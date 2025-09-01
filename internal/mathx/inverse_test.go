package mathx

import "testing"

func TestModInverse(t *testing.T) {
	cases := []struct {
		a, m, want int
		ok         bool
	}{
		{7, 26, 15, true},
		{3, 11, 4, true},
		{2, 4, 0, false},
		{10, 17, 12, true},
	}
	for _, c := range cases {
		got, ok := ModInverse(c.a, c.m)
		if ok != c.ok || (ok && got != c.want) {
			t.Errorf("ModInverse(%d,%d)=%d,%v; want %d,%v", c.a, c.m, got, ok, c.want, c.ok)
		}
	}
}
