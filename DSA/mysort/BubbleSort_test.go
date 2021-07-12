package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			[]int{22, 34, 3, 40, 18, 4},
			[]int{3, 4, 18, 22, 34, 40},
		},
	}

	for _, test := range tests {
		// 因冒泡算法是原地排序，因此额外复制一份数据，方式数据已排序影响后面的优化算法的测试
		input := make([]int, len(test.input))
		copy(input, test.input)

		BubbleSort(input)
		assert.Equal(t, test.output, input)
	}

	for _, v := range tests {
		BubbleSortV1(v.input)
		assert.Equal(t, v.output, v.input)
	}
}
