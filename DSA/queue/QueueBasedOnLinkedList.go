package queue

import (
	ll "github.com/frlute/go-playground/DSA/linkedlist"
)

// LinkedListQueue _
type LinkedListQueue struct {
	head   *ll.LinkedNode
	tail   *ll.LinkedNode
	length int
}

// NewLinkedListQueue _
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (q *LinkedListQueue) enqueue(item interface{}) {
	node := &ll.LinkedNode{Value: item, Next: nil}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.Next = node
		q.tail = node
	}

	q.length++
}

func (q *LinkedListQueue) dequeue() interface{} {
	if q.head == nil {
		return nil
	}
	val := q.head.Value
	q.head = q.head.Next
	q.length--

	return val
}
