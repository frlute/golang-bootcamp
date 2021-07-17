package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getFirstNonRepeatedChar(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			"hello",
			"h",
		},
		{
			"swiss",
			"w",
		},
		{
			"simplest",
			"i",
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := getFirstNonRepeatedChar(test.input)
			assert.Equal(t, test.output, res)

			assert.Equal(t, test.output, getFirstNonRepeatedCharV1(test.input))
		})
	}
}
