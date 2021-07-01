package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DynamicAppend(t *testing.T) {
	p := []byte{2, 3, 5}
	p = AppendByte(p, 0, 2)
	assert.Equal(t, p, []byte{2, 3, 5, 0, 2})
}
