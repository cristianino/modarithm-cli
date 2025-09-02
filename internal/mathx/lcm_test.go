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

func TestLCMMultiple(t *testing.T) {
	cases := []struct {
		numbers []int
		want    int
	}{
		{[]int{}, 0},
		{[]int{12}, 12},
		{[]int{12, 18}, 36},
		{[]int{4, 6, 8}, 24},
		{[]int{2, 3, 4, 5}, 60},
		{[]int{12, 18, 24}, 72},
		{[]int{0, 12, 18}, 0},
		{[]int{-12, 18, 24}, 72},
		{[]int{15, 20, 25}, 300},
	}
	for _, c := range cases {
		got := LCMMultiple(c.numbers...)
		if got != c.want {
			t.Errorf("LCMMultiple(%v)=%d; want %d", c.numbers, got, c.want)
		}
	}
}
