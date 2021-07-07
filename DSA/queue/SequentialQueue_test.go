package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SequentialQueue_lifecycle(t *testing.T) {
	as := assert.New(t)

	s := NewSequentialQueue(2)

	as.Equal(true, s.IsEmpty())
	as.Equal(false, s.IsFull())

	// enqueue
	s.Enqueue("hello")
	s.Enqueue(",")
	s.Enqueue("world")

	as.Equal(true, s.IsFull())
	as.Equal(false, s.IsEmpty())

	// dequeue
	as.Equal("hello", s.Dequeue())
	as.Equal(false, s.IsFull())
	as.Equal(",", s.Dequeue())
	as.Equal(true, s.IsEmpty())
	as.Nil(s.Dequeue())
}

func Test_QueueMax(t *testing.T) {
	s := newSequentialQueueArray(5)

	s.Enqueue(3)
	s.Enqueue(9)
	s.Enqueue(-1)
	s.Enqueue(8)

	assert.Equal(t, 9, s.Max())
}
