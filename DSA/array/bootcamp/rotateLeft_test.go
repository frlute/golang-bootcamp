package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotateLeft(t *testing.T) {
	tests := []struct {
		input  []int32
		d      int32
		output []int32
	}{
		{
			[]int32{1, 2, 3, 4, 5},
			4,
			[]int32{5, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := rotateLeft(test.d, test.input)
			assert.Equal(t, test.output, res)

			assert.Equal(t, test.output, rotateLeftV1(test.d, test.input))
		})
	}
}
