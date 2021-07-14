package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inputArgs struct {
	A []int
	B []int
}

func Test_intersect(t *testing.T) {
	testCases := []struct {
		args   inputArgs
		expect []int
	}{
		{
			args: inputArgs{
				A: []int{2, 3, 3, 5, 7, 11},
				B: []int{3, 3, 7, 15, 31},
			},
			expect: []int{3, 7},
		},
	}

	for _, c := range testCases {
		actual := intersect(c.args.A, c.args.B)
		assert.Equal(t, c.expect, actual)
	}
}
