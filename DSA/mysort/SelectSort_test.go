package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectSort(t *testing.T) {
	tests := []struct {
		intput []int
		output []int
	}{
		{
			[]int{4, 5, 3, 9, 1},
			[]int{1, 3, 4, 5, 9},
		},
	}

	for _, test := range tests {
		SelectSort(test.intput)
		assert.Equal(t, test.output, test.intput)
	}
}
