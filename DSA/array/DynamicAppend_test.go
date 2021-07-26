package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DynamicAppend(t *testing.T) {
	p := make([]byte, 0, 3)
	p = AppendByte(p, 1, 2)
	p = AppendByte(p, 3)
	p = AppendByte(p, 4)
	assert.Equal(t, p, []byte{1, 2, 3, 4})
}
