package binarySearch

import (
	"testing"
)

func TestFirstGTEBSearch(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		expect int
	}{
		{
			input:  []int{2, 5, 9},
			target: 13,
			expect: -1,
		},
		{
			input:  []int{5, 4, 3, 2, 1},
			target: 3,
			expect: 0,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			target: 3,
			expect: 2,
		},
	}

	for _, c := range cases {
		got := FirstGTEBSearch(c.input, c.target)
		if got != c.expect {
			t.Errorf("FirstGTEBSearch(%v, %d) == %d， want=%d", c.input, c.target, got, c.expect)
		}
	}
}
