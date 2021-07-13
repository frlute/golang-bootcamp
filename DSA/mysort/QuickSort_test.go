package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			[]int{5, 1, 1, 2, 0, 0},
			[]int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, test := range tests {
		// 因是原地排序，因此额外复制一份数据，方式数据已排序影响后面的优化算法的测试
		input := make([]int, len(test.input))
		copy(input, test.input)

		QuickSort(input)
		assert.Equal(t, test.output, input)
	}

	for _, test := range tests {
		QuickSortV1(test.input)
		assert.Equal(t, test.output, test.input)
	}
}
