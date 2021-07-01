package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	cases := []interface{}{1, 4, 5, 9, 10, 12}

	actual := Filter(cases, func(item interface{}) bool {
		if item.(int)%2 != 0 {
			return false
		}
		return true
	})
	assert.Equal(t, []interface{}{4, 10, 12}, actual)
}
