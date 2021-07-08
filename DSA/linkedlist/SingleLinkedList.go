package linkedlist

import (
	"fmt"
	"strings"
	"sync"
)

type LinkedListNode struct {
	Value interface{}
	Next  *LinkedListNode
}

// SingleLinkedList Definition for singly-linked list.
type SingleLinkedList struct {
	Head *LinkedListNode
	// 记录 tail 节点方便部分操作, 减少不断循环获取尾节点
	Tail *LinkedListNode
	// 从零开始为头结点
	length int
	sync.Mutex
}

// CreateSingleLinkedNode create single linked list instancs
func CreateSingleLinkedNode(value interface{}) LinkedList {
	return newSingleLinkedNode(value)
}

func newSingleLinkedNode(value interface{}) *SingleLinkedList {
	node := &LinkedListNode{
		Value: value,
		Next:  nil,
	}
	return &SingleLinkedList{
		Head: node,
		Tail: node,
	}
}

// Append adds an item to the end of the linked list
func (ll *SingleLinkedList) Append(value interface{}) {
	node := &LinkedListNode{
		Value: value,
		Next:  nil,
	}
	// TODO ll == nil 时需要考虑吗，目前来说 ll 不能为指针
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		ll.Tail.Next = node
		ll.Tail = node
	}

	ll.length++
}

// Prepend add an item to the begin of the linked list
func (ll *SingleLinkedList) Prepend(value interface{}) {
	node := &LinkedListNode{
		Value: value,
		Next:  nil,
	}

	head := ll.Head
	if head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		node.Next = head
		ll.Head = node
	}

	ll.length++
}

// Insert add an item to the specific position in linked list
func (ll *SingleLinkedList) Insert(value interface{}, pos int) error {
	if ll.length < pos {
		return fmt.Errorf("pos %d it out of the linked list %d", pos, ll.length)
	} else if ll.Size() == pos {
		ll.Append(value)
	} else if pos == 0 {
		ll.Prepend(value)
	} else {
		node := &LinkedListNode{
			Value: value,
			Next:  nil,
		}

		head := ll.Head
		var preNode *LinkedListNode
		for i := 0; i < ll.length; i++ {
			if i == pos {
				node.Next = head
				preNode.Next = node
				break
			}
			preNode = head
			head = head.Next
		}
		ll.length++
	}

	return nil
}

// RemoveAt remove the item at the specified position item in linked list
func (ll *SingleLinkedList) RemoveAt(pos int) {
	if ll.length <= pos {
		fmt.Printf("pos %d it out of the linked list %d\n", pos, ll.length)
		return
	}

	head := ll.Head
	var preNode *LinkedListNode
	for i := 0; i < ll.length; i++ {
		if i == pos {
			// 删除的位置是尾节点是，前置节点指为 nil 即可
			if head.Next == nil {
				ll.Tail = preNode
				// 只有一个节点时
				if preNode == nil {
					ll.Head = nil
					break
				}
			}
			preNode.Next = head.Next

			break
		}
		preNode = head
		head = head.Next
	}
	ll.length--
}

// Delete delete the specific item in linked list
func (ll *SingleLinkedList) Delete(value interface{}) {
	head := ll.Head
	var preNode *LinkedListNode
	// TODO 待用 for head.Next != nil 优化
	for pos := 0; pos < ll.length; pos++ {
		if head.Value == value {
			// 尾节点时特殊处理
			if head.Next == nil {
				if preNode != nil {
					preNode.Next = head.Next
				} else {
					// 只有一个节点时
					ll.Head = nil
				}
				ll.Tail = preNode
			} else {
				preNode.Next = head.Next
			}
			ll.length--
			break
		}
		preNode = head
		head = head.Next
	}
}

// IndexOf get the value position in linked list
func (ll *SingleLinkedList) IndexOf(value interface{}) int {

	head := ll.Head
	for pos := 0; pos < ll.length; pos++ {
		if head.Value == value {
			return pos
		}
		head = head.Next
	}
	return -1
}

// IsEmpty judge the linked list is empty
func (ll *SingleLinkedList) IsEmpty() bool {
	if ll == nil || ll.Head == nil {
		return true
	}
	return false
}

// Size get the linked list length
func (ll *SingleLinkedList) Size() int {
	if ll.Head == nil {
		return 0
	}
	return ll.length
}

// Display a string representation of the list
// TODO 如果是循环链表时，disply 和 string 目前会无限循环，需要改为按 pos 遍历
func (ll *SingleLinkedList) Display() {
	if ll.Head == nil {
		return
	}
	fmt.Printf("Display Single Linked List: ")

	head := ll.Head
	for head != nil {
		if head.Next == nil {
			fmt.Printf("%v\n", head.Value)
		} else {
			fmt.Printf("%v->", head.Value)
		}
		head = head.Next
	}
}

// String a string representation of the list
func (ll *SingleLinkedList) String() string {
	if ll.Head == nil {
		return "This is a empty single linked list."
	}
	var output strings.Builder
	output.WriteString("Single Linked List: ")

	head := ll.Head
	for head != nil {
		if head.Next == nil {
			output.WriteString(fmt.Sprintf("%v\n", head.Value))
		} else {
			output.WriteString(fmt.Sprintf("%v->", head.Value))
		}
		head = head.Next
	}
	return output.String()
}

// Reset reset all linked list
func (ll *SingleLinkedList) Reset() {
	ll.Head = nil
	ll.Tail = nil
	ll.length = 0
}
