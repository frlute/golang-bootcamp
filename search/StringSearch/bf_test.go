package StringSearch

import "testing"

func TestBf(t *testing.T) {
	cases := []struct {
		src    string
		target string
		expect int
	}{
		{
			"abcd227fac",
			"ac",
			8,
		},
	}

	for _, c := range cases {
		got := bf(c.src, c.target)
		if got != c.expect {
			t.Errorf("bf(%v, %v) = %d, want=%d", c.src, c.target, got, c.expect)
		}
	}
}
