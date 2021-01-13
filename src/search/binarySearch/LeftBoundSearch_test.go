package binarySearch

import (
	"testing"
)

func TestLeftBoundSearch(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		expect int
	}{
		{
			input:  []int{},
			target: 2,
			expect: -1,
		},
		{
			input:  []int{2, 5, 9},
			target: 13,
			expect: -1,
		},
		{
			input:  []int{1, 2, 2, 2, 3},
			target: 2,
			expect: 1,
		},
	}

	for _, c := range cases {
		got := LeftBoundSearch(c.input, c.target)
		if got != c.expect {
			t.Errorf("LeftBoundSearch(%v, %d) == %d， want=%d", c.input, c.target, got, c.expect)
		}
		got1 := LeftBoundSearchV1(c.input, c.target)
		if got1 != c.expect {
			t.Errorf("LeftBoundSearchV1(%v, %d) == %d， want=%d", c.input, c.target, got1, c.expect)
		}
	}
}
