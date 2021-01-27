package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStack(t *testing.T) {
	s := NewArrayStack(2)
	as := assert.New(t)

	// push
	as.Equal(true, s.push("hello"))
	as.Equal(true, s.push(","))
	as.Equal(false, s.push("world"))

	// pop
	as.Equal(",", s.pop())
	as.Equal("hello", s.pop())
	as.Equal("", s.pop())
}
