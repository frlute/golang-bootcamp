package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueueBasedArray(t *testing.T) {
	s := NewQueueBasedArray(2)
	as := assert.New(t)

	// enqueue
	as.Equal(true, s.Enqueue("hello"))
	as.Equal(true, s.Enqueue(","))
	as.Equal(false, s.Enqueue("world"))

	// dequeue
	as.Equal("hello", s.Dequeue())
	as.Equal(",", s.Dequeue())
	as.Equal("", s.Dequeue())

	as.Equal(true, s.Enqueue("test"))
}

func Test_QueueMax(t *testing.T) {
	s := newQueueBasedArrayArray(5)

	s.Enqueue(3)
	s.Enqueue(9)
	s.Enqueue(-1)
	s.Enqueue(8)

	assert.Equal(t, 9, s.Max())
}
