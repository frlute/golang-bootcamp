package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueBasedOnLinkedList(t *testing.T) {
	q := NewLinkedListQueue()
	as := assert.New(t)

	// enqueue
	q.enqueue("hello")
	q.enqueue("world")
	as.Equal(2, q.length)

	// // dequeue
	as.Equal("hello", q.dequeue())
	as.Equal("world", q.dequeue())
	as.Equal(0, q.length)
}
