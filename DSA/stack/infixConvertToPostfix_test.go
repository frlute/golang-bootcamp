package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_infixConvertToPostfix(t *testing.T) {
	tests := []struct {
		input  string
		output string
		eval   int
	}{
		{
			"5+6",
			"56+",
			11,
		},
		{
			"2+3*4",
			"234*+",
			14,
		},
		{
			"2*(6+2*4)*5-8/4",
			"2624*+*5*84/-",
			138,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			infix := newInfix(test.input)
			res := infix.convertToPostfix()
			assert.Equal(t, test.output, res)

			calcValue := postfixEvaluation(res)
			assert.Equal(t, test.eval, calcValue)
		})
	}
}
