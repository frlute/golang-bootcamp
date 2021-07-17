package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			"hello",
			"olleh",
		},
		{
			"text",
			"txet",
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := reverse(test.input)
			assert.Equal(t, test.output, res)

			assert.Equal(t, test.output, reverseV1(test.input))
		})
	}
}
