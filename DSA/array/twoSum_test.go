package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		output []int
	}{
		{
			input:  []int{1},
			target: 5,
			output: nil,
		},
		{
			input:  []int{1, 2, 5, 6, 8, 10},
			target: 13,
			output: []int{2, 4},
		},
	}

	for _, item := range testCases {
		t.Run("test two sum", func(t *testing.T) {
			actual := twoSum(item.input, item.target)
			assert.Equal(t, item.output, actual)

			assert.Equal(t, item.output, twoSumBasedHash(item.input, item.target))

			// 此解法数组必须已排序
			assert.Equal(t, item.output, twoSumBasedInvariants(item.input, item.target))
		})
	}
}
