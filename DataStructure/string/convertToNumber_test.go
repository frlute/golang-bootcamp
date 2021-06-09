package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertToInt(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{
			input:  "12345",
			output: 12345,
		},
		{
			input:  "-12",
			output: -12,
		},
	}

	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			assert.Equal(t, convertToInt(c.input), c.output)
			assert.Equal(t, convertToIntV2(c.input), c.output)
		})
	}
}
