package binarySearch

import (
	"testing"
)

func TestLastLeBSearch(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		expect int
	}{
		{
			input:  []int{15, 19, 20},
			target: 13,
			expect: -1,
		},
		{
			input:  []int{5, 4, 3, 2, 1},
			target: 3,
			expect: 4,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			target: 3,
			expect: 2,
		},
	}

	for _, c := range cases {
		got := LastLeBSearch(c.input, c.target)
		if got != c.expect {
			t.Errorf("LastLeBSearch(%v, %d) == %d， want=%d", c.input, c.target, got, c.expect)
		}
	}
}
