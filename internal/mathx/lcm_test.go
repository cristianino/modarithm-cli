package mathx

import "testing"

func TestLCM(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{12, 18, 36},
		{0, 18, 0},
		{12, 0, 0},
		{-12, 18, 36},
		{18, -12, 36},
	}
	for _, c := range cases {
		got := LCM(c.a, c.b)
		if got != c.want {
			t.Errorf("LCM(%d,%d)=%d; want %d", c.a, c.b, got, c.want)
		}
	}
}
