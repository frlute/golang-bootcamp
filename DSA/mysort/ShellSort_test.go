package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShellSort(t *testing.T) {
	tests := []struct {
		intput []int
		output []int
	}{
		{
			[]int{4, 5, 3, 9, 1},
			[]int{1, 3, 4, 5, 9},
		},
		{
			[]int{11, 9, 7, 5, 4, 2, 1, -1},
			[]int{-1, 1, 2, 4, 5, 7, 9, 11},
		},
	}

	for _, test := range tests {
		ShellSort(test.intput)
		assert.Equal(t, test.output, test.intput)
	}
}
