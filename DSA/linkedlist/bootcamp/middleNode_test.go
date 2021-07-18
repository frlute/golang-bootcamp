package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{3, 4, 5},
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6},
			output: []int{4, 5, 6},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := runIntSingleLinkedListFunc(test.input, test.output, middleNode)

			assert.Equal(t, res.expect.String(), res.result.String())

			res1 := runIntSingleLinkedListFunc(test.input, test.output, middleNodeV1)
			assert.Equal(t, res1.expect.String(), res1.result.String())
		})
	}
}
