package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindromic(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{
			"hello",
			false,
		},
		{
			"hellolleh",
			true,
		},
		{
			"heeh",
			true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			res := isPalindromic(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}
}
