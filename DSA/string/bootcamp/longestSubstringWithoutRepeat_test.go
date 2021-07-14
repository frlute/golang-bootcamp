package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestSubstringWithoutRepeat(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{
			input:  "abcabcbb",
			output: 3,
		},
		{
			input:  "bbbbb",
			output: 1,
		},
		{
			input:  "pwwkew",
			output: 3,
		},
		{
			input:  "",
			output: 0,
		},
	}

	for _, item := range cases {
		t.Run("", func(t *testing.T) {
			res := longestSubstringWithoutRepeat(item.input)
			assert.Equal(t, item.output, res)
		})
	}
}
