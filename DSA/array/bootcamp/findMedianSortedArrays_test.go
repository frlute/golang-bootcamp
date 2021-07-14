package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type Args struct {
		input1 []int
		input2 []int
	}

	tests := []struct {
		args   Args
		output float64
	}{
		{
			args: Args{
				input1: []int{1, 3},
				input2: []int{2},
			},
			output: 2.0000,
		},
		{
			args: Args{
				input1: []int{1, 3},
				input2: []int{2, 4},
			},
			output: 2.50000,
		},
		{
			args: Args{
				input1: []int{0, 0},
				input2: []int{0, 0},
			},
			output: 0.00000,
		},
		{
			args: Args{
				input1: []int{},
				input2: []int{1},
			},
			output: 1.00000,
		},
		{
			args: Args{
				input1: []int{2},
				input2: []int{},
			},
			output: 2.00000,
		},
		{
			args: Args{
				input1: []int{4, 5, 9},
				input2: []int{2, 4, 7, 11},
			},
			output: 5.00000,
		},
	}
	for _, test := range tests {
		res := findMedianSortedArrays(test.args.input1, test.args.input2)
		assert.Equal(t, test.output, res)
	}
}
