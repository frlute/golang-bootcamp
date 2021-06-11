package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StackBasedArray(t *testing.T) {
	s := NewArrayStack(2, false)
	as := assert.New(t)

	// push
	as.Equal(true, s.Push("hello"))
	as.Equal(true, s.Push(","))
	as.Equal(false, s.Push("world"))

	// pop
	as.Equal(",", s.Pop())
	as.Equal("hello", s.Pop())
	as.Equal("", s.Pop())
}
