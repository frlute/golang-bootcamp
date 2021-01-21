package binarySearch

import (
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		input  []int
		target int
		expect int
	}{
		{
			[]int{1, 2, 4, 5, 7},
			4,
			2,
		},
	}

	for _, c := range cases {
		got := Search(c.input, c.target)
		if got != c.expect {
			t.Errorf("Search(%v, %d) == %d, want=%d", c.input, c.target, got, c.expect)
		}
	}

	for _, c := range cases {
		got := SearchByRecursion(c.input, c.target)
		if got != c.expect {
			t.Errorf("SearchByRecursion(%v, %d) == %d, want=%d", c.input, c.target, got, c.expect)
		}
	}
}
