package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayQueue(t *testing.T) {
	s := NewArrayQueue(2)
	as := assert.New(t)

	// enqueue
	as.Equal(true, s.enqueue("hello"))
	as.Equal(true, s.enqueue(","))
	as.Equal(false, s.enqueue("world"))

	// dequeue
	as.Equal("hello", s.dequeue())
	as.Equal(",", s.dequeue())
	as.Equal("", s.dequeue())

	as.Equal(true, s.enqueue("test"))
}
