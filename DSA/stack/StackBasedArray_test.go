package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StackBasedArray_lifecycle(t *testing.T) {
	s := NewStackBasedArray(2)
	as := assert.New(t)

	as.Equal(true, s.IsEmpty())

	// push
	s.Push(1)
	s.Push(2)

	as.Equal(true, s.IsFull())
	as.Equal(2, s.Peek())
	as.Equal(true, s.IsFull())

	as.Equal(2, s.Pop())
	as.Equal(false, s.IsFull())
	as.Equal(1, s.Pop())
	as.Equal(true, s.IsEmpty())
}
