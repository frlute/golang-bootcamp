package morestrings

import (
	"testing"
)

func TestReverseRuns(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRuns(%q) == %q， want=%q", c.in, got, c.want)
		}
	}
}
