package mathx

import (
	"testing"
)

func TestMultiplesSym(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		limit   int
		want    Multiples
		wantErr bool
	}{
		{
			name:  "basic symmetric multiples",
			n:     5,
			limit: 2,
			want:  Multiples{Negatives: []int{-10, -5}, Zero: 0, Positives: []int{5, 10}},
		},
		{
			name:  "zero limit",
			n:     7,
			limit: 0,
			want:  Multiples{Negatives: []int{}, Zero: 0, Positives: []int{}},
		},
		{
			name:    "negative limit error",
			n:       3,
			limit:   -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MultiplesSym(tt.n, tt.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiplesSym() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !equalMultiples(got, tt.want) {
				t.Errorf("MultiplesSym() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplesAsym(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		limitNeg int
		limitPos int
		want     Multiples
		wantErr  bool
	}{
		{
			name:     "basic asymmetric multiples",
			n:        3,
			limitNeg: 1,
			limitPos: 2,
			want:     Multiples{Negatives: []int{-3}, Zero: 0, Positives: []int{3, 6}},
		},
		{
			name:     "negative limitNeg error",
			n:        5,
			limitNeg: -1,
			limitPos: 2,
			wantErr:  true,
		},
		{
			name:     "negative limitPos error",
			n:        5,
			limitNeg: 1,
			limitPos: -2,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MultiplesAsym(tt.n, tt.limitNeg, tt.limitPos)
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiplesAsym() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !equalMultiples(got, tt.want) {
				t.Errorf("MultiplesAsym() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiples_RenderGroup(t *testing.T) {
	m := Multiples{
		Negatives: []int{-10, -5},
		Zero:      0,
		Positives: []int{5, 10},
	}

	tests := []struct {
		name  string
		group Group
		want  []int
	}{
		{"negative group", GroupNeg, []int{-10, -5}},
		{"zero group", GroupZero, []int{0}},
		{"positive group", GroupPos, []int{5, 10}},
		{"all group", GroupAll, []int{-10, -5, 0, 5, 10}},
		{"invalid group", "invalid", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := m.RenderGroup(tt.group)
			if !equalSlice(got, tt.want) {
				t.Errorf("RenderGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalMultiples(a, b Multiples) bool {
	return equalSlice(a.Negatives, b.Negatives) &&
		a.Zero == b.Zero &&
		equalSlice(a.Positives, b.Positives)
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
