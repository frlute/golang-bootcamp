package stack

import (
	"fmt"
	"sync"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

type StackBasedLinkedList struct {
	top      *ll.LinkedListNode
	length   int64 // 栈中元素长度，从 0 开始
	capacity int64 // 栈容量
	sync.Mutex
}

// NewStackBasedLinkedList _
func NewStackBasedLinkedList(capacity int64) Stack {
	if capacity < 0 {
		capacity = 0
	}
	if capacity == 0 {
		capacity = 10
	}

	return &StackBasedLinkedList{
		top:      &ll.LinkedListNode{},
		length:   0,
		capacity: capacity,
	}
}

func (s *StackBasedLinkedList) Push(item interface{}) {
	if s.IsFull() {
		fmt.Println("Could not insert data, Stack is full.")
		return
	}
	node := &ll.LinkedListNode{
		Value: item,
		Next:  nil,
	}
	node.Next = s.top
	s.top = node
	s.length++
}

// Pop Removing an element from the stack.
func (s *StackBasedLinkedList) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("Stack is empty.")
		return nil
	}
	value := s.top.Value
	s.top = s.top.Next
	s.length--
	return value
}

// Peek get the top data element of the stack, without removing it.
func (s *StackBasedLinkedList) Peek() interface{} {
	return s.top.Value
}

// IsFull check if stack is full.
func (s *StackBasedLinkedList) IsFull() bool {
	return s.length == s.capacity
}

// IsEmpty check if stack is empty.
func (s *StackBasedLinkedList) IsEmpty() bool {
	return s.length == 0
}

// Flush _
func (s *StackBasedLinkedList) Flush() {
	s.top = nil
	s.length = 0
}

// Display _
func (s *StackBasedLinkedList) Display() {
	list := &ll.SingleLinkedList{
		Head: s.top,
	}
	list.Display()
}
