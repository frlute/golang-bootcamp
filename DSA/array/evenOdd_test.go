package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_even_odd(t *testing.T) {
	assert.Equal(t, []int{2, 4, 8, 3, 9, 5, 1}, evenOdd([]int{1, 5, 8, 3, 4, 9, 2}))
}
