package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ChainQueue_lifecycle(t *testing.T) {
	as := assert.New(t)

	q := NewChainQueue(3)

	as.Equal(true, q.IsEmpty())
	as.Equal(false, q.IsFull())

	// enqueue
	q.Enqueue("hello")
	q.Enqueue("world")
	as.Equal(false, q.IsEmpty())
	as.Equal(false, q.IsFull())
	q.Enqueue("!")
	as.Equal(true, q.IsFull())
	q.Enqueue("?")

	// dequeue
	as.Equal("hello", q.Dequeue())
	as.Equal("world", q.Dequeue())
	as.Equal(false, q.IsEmpty())
	as.Equal(false, q.IsFull())
	as.Equal("!", q.Dequeue())
	as.Nil(q.Dequeue())
	as.Equal(true, q.IsEmpty())
}
