package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueBasedOnLinkedList(t *testing.T) {
	s := NewLinkedListQueue()
	as := assert.New(t)

	// enqueue
	as.Equal(true, s.enqueue("hello"))
	as.Equal(true, s.enqueue("world"))

	// dequeue
	as.Equal("hello", s.dequeue())
	as.Equal("world", s.dequeue())

	as.Equal(true, s.enqueue("test"))
}
