package binarySearch

import (
	"testing"
)

func TestRightBoundSearch(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		expect int
	}{
		{
			input:  []int{1, 2, 2, 2, 3},
			target: 2,
			expect: 3,
		},
	}

	for _, c := range cases {
		got := RightBoundSearch(c.input, c.target)
		if got != c.expect {
			t.Errorf("RightBoundSearch(%v, %d) == %d， want=%d", c.input, c.target, got, c.expect)
		}
	}
}
