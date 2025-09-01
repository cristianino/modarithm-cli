package mathx

import "testing"

func TestGCD(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{84, 30, 6},
		{0, 0, 0},
		{12, 0, 12},
		{0, 18, 18},
		{-12, 18, 6},
		{18, -12, 6},
	}
	for _, c := range cases {
		got := GCD(c.a, c.b)
		if got != c.want {
			t.Errorf("GCD(%d,%d)=%d; want %d", c.a, c.b, got, c.want)
		}
	}
}
