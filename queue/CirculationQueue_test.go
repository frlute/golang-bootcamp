package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCirculationQueue(t *testing.T) {
	q := NewCirculationQueue(2)
	as := assert.New(t)

	// enqueue
	as.Equal(true, q.enqueue("hello"))
	// 已经满了，会浪费一个位置
	as.Equal(false, q.enqueue("world"))

	// dequeue
	as.Equal("hello", q.dequeue())
	as.Equal(nil, q.dequeue())
}
