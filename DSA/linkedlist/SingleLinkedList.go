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
	// TODO 此处可以改为使用哨兵节点作为头结点
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

	if ll.IsEmpty() {
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
	if ll.Size() < pos {
		return fmt.Errorf("pos %d it out of the linked list %d", pos, ll.Size())
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
		// TODO 也可以用哨兵节点方式，省去记录前置节点的过程
		var preNode *LinkedListNode
		for i := 0; i < ll.Size(); i++ {
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

// RemoveAt remove the item at the specified position item in linked list, pos 从 0 开始
func (ll *SingleLinkedList) RemoveAt(pos int) {
	if ll.Size() <= pos {
		fmt.Printf("pos %d it out of the linked list %d\n", pos, ll.Size())
		return
	}

	head := ll.Head
	var preNode *LinkedListNode
	for i := 0; i < ll.Size(); i++ {
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
	// var preNode *LinkedListNode
	// TODO 待用 for head.Next != nil 优化
	// for pos := 0; pos < ll.Size(); pos++ {
	// 	if head.Value == value {
	// 		// 尾节点时特殊处理
	// 		if head.Next == nil {
	// 			if preNode != nil {
	// 				preNode.Next = head.Next
	// 			} else {
	// 				// 只有一个节点时
	// 				ll.Head = nil
	// 			}
	// 			ll.Tail = preNode
	// 		} else {
	// 			preNode.Next = head.Next
	// 		}
	// 		ll.length--
	// 		break
	// 	}
	// 	preNode = head
	// 	head = head.Next
	// }

	// 利用哨兵节点简化业务
	dummy := &LinkedListNode{}
	dummy.Next = head
	res := dummy
	for dummy.Next != nil {
		if dummy.Next.Value == value {
			next := dummy.Next.Next
			dummy.Next = next
			ll.length--
			ll.Head = res.Next
			// 尾节点时更新新的尾节点
			if next == nil {
				// 只有一个节点时，删除完 Tail 就是 nil 了，而非空节点
				if dummy.Value == nil {
					ll.Tail = nil
				} else {
					ll.Tail = dummy
				}
			}
			return
		}
		dummy = dummy.Next
	}
}

// IndexOf get the value position in linked list
func (ll *SingleLinkedList) IndexOf(value interface{}) int {

	head := ll.Head
	for pos := 0; pos < ll.Size(); pos++ {
		if head.Value == value {
			return pos
		}
		head = head.Next
	}
	return -1
}

// Search get the value node and position in linked list
func (ll *SingleLinkedList) Search(value interface{}) interface{} {

	head := ll.Head
	for pos := 0; pos < ll.Size(); pos++ {
		if head.Value == value {
			return head
		}
		head = head.Next
	}
	return nil
}

func (ll *SingleLinkedList) Index(index int) *LinkedListNode {

	head := ll.Head
	for pos := 0; pos < ll.Size(); pos++ {
		if pos == index {
			return head
		}
		head = head.Next
	}
	return nil
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
	// 如果 length 未正确记录时
	if ll.Head != nil && ll.length == 0 {
		cur := ll.Head
		// TODO 如果链表成环此种方式会造成死循环
		for cur != nil {
			ll.length++
			cur = cur.Next
		}
	}

	return ll.length
}

// Display a string representation of the list
func (ll *SingleLinkedList) Display() {
	if ll.IsEmpty() {
		return
	}
	fmt.Printf("Display Single Linked List: ")

	head := ll.Head
	for i := 0; i < ll.Size(); i++ {
		// 兼容链表成环
		if head.Next == nil || i == ll.Size() {
			fmt.Printf("%v\n", head.Value)
		} else {
			fmt.Printf("%v->", head.Value)
		}
		head = head.Next
	}
}

// String a string representation of the list
func (ll *SingleLinkedList) String() string {
	if ll.IsEmpty() {
		fmt.Println("[info] This is a empty single linked list.")
		return ""
	}

	var output strings.Builder
	output.WriteString("Single Linked List: ")

	head := ll.Head
	for i := 0; i < ll.Size(); i++ {
		// 兼容链表成环
		if head.Next == nil || i == ll.Size() {
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
	if ll.IsEmpty() {
		return
	}
	ll.Head = nil
	ll.Tail = nil
	ll.length = 0
}

// AppendIntoTail _
// TODO 方便构造循环成环数据, 暂未考虑边界条件
func (ll *SingleLinkedList) AppendIntoTail(node *LinkedListNode) {
	if node == nil {
		return
	}
	ll.length++
	ll.Tail.Next = node
}
