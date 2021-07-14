package bootcamp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3},
			output: []int{1, 2, 4},
		},
		{
			input:  []int{4, 3, 2, 1},
			output: []int{4, 3, 2, 2},
		},
		{
			input:  []int{0},
			output: []int{1},
		},
		{
			input:  []int{9},
			output: []int{1, 0},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%+v plus 1", test.input)
		t.Run(name, func(t *testing.T) {
			res := plusOne(test.input)
			assert.Equal(t, test.output, res)
		})
	}
}
