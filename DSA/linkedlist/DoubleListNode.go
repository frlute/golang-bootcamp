package linkedlist

import (
	"sync"
)

type DoubleLinkedListNode struct {
	Value interface{}
	Prev  *DoubleLinkedListNode
	Next  *DoubleLinkedListNode
}

// DoubleLinkedList _
type DoubleLinkedList struct {
	head *DoubleLinkedListNode
	tail *DoubleLinkedListNode
	// 从零开始为头结点
	length int
	sync.Mutex
}

// CreateDoubleLinkedList create single linked list instancs
func CreateDoubleLinkedList(value interface{}) LinkedList {
	return newDoubleLinkedList(value)
}

func newDoubleLinkedList(value interface{}) *DoubleLinkedList {
	newNode := &DoubleLinkedListNode{
		Value: value,
		Prev:  nil,
		Next:  nil,
	}
	return &DoubleLinkedList{
		head: newNode,
		tail: newNode,
	}
}

// Append adds an item to the end of the linked list
func (dl *DoubleLinkedList) Append(value interface{}) {

}

// Prepend add an item to the begin of the linked list
func (dl *DoubleLinkedList) Prepend(value interface{}) {}

// Insert add an item to the specific position in linked list
func (dl *DoubleLinkedList) Insert(value interface{}, pos int) error {
	return nil
}

// RemoveAt remove the item at the specified position item in linked list
func (dl *DoubleLinkedList) RemoveAt(pos int) {

}

// Delete delete the specific item in linked list
func (dl *DoubleLinkedList) Delete(value interface{}) {

}

// IndexOf get the value position in linked list
func (dl *DoubleLinkedList) IndexOf(value interface{}) int {
	return -1
}

// IsEmpty judge the linked list is empty
func (dl *DoubleLinkedList) IsEmpty() bool {
	return false
}

// Size get the linked list length
func (dl *DoubleLinkedList) Size() int {
	return dl.length
}

// Display a string representation of the list
func (dl *DoubleLinkedList) Display() {
}

// String a string representation of the list
func (dl *DoubleLinkedList) String() string {
	return ""
}

// Reset reset adl linked list
func (dl *DoubleLinkedList) Reset() {
	dl.head = nil
	dl.tail = nil
	dl.length = 0
}
