package bootcamp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeIntervals(t *testing.T) {
	tests := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			output: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			input: [][]int{
				{1, 4},
				{4, 5},
			},
			output: [][]int{
				{1, 5},
			},
		},
		{
			input: [][]int{
				{1, 4},
				{0, 5},
			},
			output: [][]int{
				{0, 5},
			},
		},
		{
			input: [][]int{
				{2, 3},
				{4, 5},
				{6, 7},
				{8, 9},
				{1, 10},
			},
			output: [][]int{
				{1, 10},
			},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("mergeIntervalsV1(%+v)", test.input)
		t.Run(name, func(t *testing.T) {
			res := mergeIntervalsV1(test.input)
			assert.Equal(t, test.output, res)
		})
	}
}

func Test_xx(t *testing.T) {
	test := [][]int{
		{2, 3},
		{4, 5},
		{6, 7},
		{8, 9},
		{1, 10},
	}
	mergeIntervals(test)
}
