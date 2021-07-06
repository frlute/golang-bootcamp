package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StackBasedLinkedList_lifecycle(t *testing.T) {
	as := assert.New(t)

	s := NewStackBasedLinkedList(2)

	as.Equal(true, s.IsEmpty())

	s.Push(1)
	s.Push(2)

	as.Equal(true, s.IsFull())
	as.Equal(2, s.Peek())
	as.Equal(true, s.IsFull())

	s.Display()

	as.Equal(2, s.Pop())
	as.Equal(false, s.IsFull())
	as.Equal(1, s.Pop())
	as.Equal(true, s.IsEmpty())

	s.Display()
}
