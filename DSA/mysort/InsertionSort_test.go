package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		intput []int
		output []int
	}{
		{
			[]int{4, 2, 1, 3},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{-1, 5, 3, 4, 0},
			[]int{-1, 0, 3, 4, 5},
		},
		{
			[]int{4, 5, 3, 9, 1},
			[]int{1, 3, 4, 5, 9},
		},
	}

	for _, test := range tests {
		assert.NotEqual(t, test.output, test.intput)
		InsertionSort(test.intput)
		assert.Equal(t, test.output, test.intput)
	}
}
