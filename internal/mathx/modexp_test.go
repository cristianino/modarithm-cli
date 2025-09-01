package mathx

import "testing"

func TestModExp(t *testing.T) {
	cases := []struct{ base, exp, mod, want int }{
		{5, 117, 19, 1},
		{2, 10, 1000, 24},
		{3, 0, 7, 1},
		{10, 1, 7, 3},
		{2, 20, 17, 16},
	}
	for _, c := range cases {
		got := ModExp(c.base, c.exp, c.mod)
		if got != c.want {
			t.Errorf("ModExp(%d,%d,%d)=%d; want %d", c.base, c.exp, c.mod, got, c.want)
		}
	}
}
