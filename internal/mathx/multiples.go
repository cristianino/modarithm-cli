package mathx

import (
	"errors"
	"slices"
)

type Group string

const (
	GroupNeg  Group = "neg"
	GroupZero Group = "zero"
	GroupPos  Group = "pos"
	GroupAll  Group = "all"
)

type Multiples struct {
	Negatives []int `json:"negatives"`
	Zero      int   `json:"zero"`
	Positives []int `json:"positives"`
}

// MultiplesSym returns symmetric multiples with validation
func MultiplesSym(n, limit int) (Multiples, error) {
	if limit < 0 {
		return Multiples{}, errors.New("limit must be >= 0")
	}

	neg := make([]int, 0, limit)
	pos := make([]int, 0, limit)

	for k := 1; k <= limit; k++ {
		neg = append(neg, -n*k)
		pos = append(pos, n*k)
	}

	// Sort negatives in ascending order: [-20, -15, -10, -5]
	slices.Sort(neg)

	return Multiples{Negatives: neg, Zero: 0, Positives: pos}, nil
}

// MultiplesAsym returns asymmetric multiples with validation
func MultiplesAsym(n, limitNeg, limitPos int) (Multiples, error) {
	if limitNeg < 0 || limitPos < 0 {
		return Multiples{}, errors.New("limits must be >= 0")
	}

	neg := make([]int, 0, limitNeg)
	pos := make([]int, 0, limitPos)

	for k := 1; k <= limitNeg; k++ {
		neg = append(neg, -n*k)
	}
	for k := 1; k <= limitPos; k++ {
		pos = append(pos, n*k)
	}

	slices.Sort(neg)

	return Multiples{Negatives: neg, Zero: 0, Positives: pos}, nil
}

// RenderGroup returns the slice according to the requested group
// For GroupAll, returns a single ordered slice: neg..., 0, pos...
func (m Multiples) RenderGroup(g Group) []int {
	switch g {
	case GroupNeg:
		return append([]int(nil), m.Negatives...)
	case GroupZero:
		return []int{0}
	case GroupPos:
		return append([]int(nil), m.Positives...)
	case GroupAll:
		all := make([]int, 0, len(m.Negatives)+1+len(m.Positives))
		all = append(all, m.Negatives...)
		all = append(all, 0)
		all = append(all, m.Positives...)
		return all
	default:
		return nil
	}
}
