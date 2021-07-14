package bootcamp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortColors(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{2, 0, 2, 1, 1, 0},
			output: []int{0, 0, 1, 1, 2, 2},
		},
		{
			input:  []int{2, 0, 1},
			output: []int{0, 1, 2},
		},
		{
			input:  []int{0},
			output: []int{0},
		},
		{
			input:  []int{1},
			output: []int{1},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("test sortColors(%+v)", test.input)
		t.Run(name, func(t *testing.T) {
			res := sortColors(test.input)
			as.Equal(test.output, res)
		})
	}
}
