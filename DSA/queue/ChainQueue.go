package queue

import (
	"fmt"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

// chainQueue _
type chainQueue struct {
	head   *ll.LinkedListNode
	tail   *ll.LinkedListNode
	length int
	// 虽然链表无需长度限制的，但是 queue 是有限的
	capacity int
}

// NewChainQueue _
func NewChainQueue(cap int) Queue {
	return &chainQueue{nil, nil, 0, cap}
}

func (q *chainQueue) Enqueue(item interface{}) {
	if q.IsFull() {
		fmt.Println("Could not insert data, Queue is full.")
		return
	}
	node := &ll.LinkedListNode{Value: item, Next: nil}
	if q.IsEmpty() {
		q.head = node
		q.tail = node
	} else {
		q.tail.Next = node
		q.tail = node
	}

	q.length++
}

func (q *chainQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	val := q.head.Value
	q.head = q.head.Next
	q.length--

	return val
}

func (q *chainQueue) Peek() interface{} {
	return q.head.Value
}

func (q *chainQueue) IsFull() bool {
	return q.length == q.capacity
}

func (q *chainQueue) IsEmpty() bool {
	return q.head == nil
}
