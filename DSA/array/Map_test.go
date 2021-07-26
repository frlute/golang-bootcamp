package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	square := func(x int) int {
		return x * x
	}

	nums := []int{1, 2, 3, 4}
	assert.Equal(t, []interface{}{1, 4, 9, 16}, Map(nums, square))
}
