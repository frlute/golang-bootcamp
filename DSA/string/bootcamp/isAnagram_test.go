package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		inputA string
		inputB string
		output bool
	}{
		{
			"word",
			"wrdo",
			true,
		},
		{
			"mary",
			"army",
			true,
		},
		{
			"tops",
			"stop",
			true,
		},
		{
			"fill",
			"fil",
			false,
		},
		{
			"bbbb",
			"b",
			false,
		},
		{
			"a",
			"a",
			true,
		},
		{
			"pure",
			"in",
			false,
		},
		{
			"sleep",
			"slep",
			false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := isAnagram(test.inputA, test.inputB)
			assert.Equal(t, test.output, res)

			assert.Equal(t, test.output, isAnagramV1(test.inputA, test.inputB))

			assert.Equal(t, test.output, isAnagramV2(test.inputA, test.inputB))
		})
	}
}
