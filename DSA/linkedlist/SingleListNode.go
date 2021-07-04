package linkedlist

import (
	"fmt"
	"strings"
	"sync"

	"github.com/frlute/golang-bootcamp/DSA/utils"
)

// SingleLinkedNode Definition for singly-linked list.
type SingleLinkedNode struct {
	head *utils.Node
	// 记录 tail 节点方便部分操作
	tail *utils.Node
	// 从零开始为头结点
	length int
	sync.Mutex
}

// CreateSingleLinkedNode create single linked list instancs
func CreateSingleLinkedNode(value interface{}) LinkedList {
	return newSingleLinkedNode(value)
}

func newSingleLinkedNode(value interface{}) *SingleLinkedNode {
	node := &utils.Node{
		Value: value,
		Next:  nil,
	}
	return &SingleLinkedNode{
		head: node,
		tail: node,
	}
}

// Append adds an item to the end of the linked list
func (ll *SingleLinkedNode) Append(value interface{}) {
	node := &utils.Node{
		Value: value,
		Next:  nil,
	}
	// TODO ll == nil 时需要考虑吗，目前来说 ll 不能为指针
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.Next = node
		ll.tail = node
	}

	ll.length++
}

// Prepend add an item to the begin of the linked list
func (ll *SingleLinkedNode) Prepend(value interface{}) {
	node := &utils.Node{
		Value: value,
		Next:  nil,
	}

	head := ll.head
	if head == nil {
		ll.head = node
		ll.tail = node
	} else {
		node.Next = head
		ll.head = node
	}

	ll.length++
}

// Insert add an item to the specific position in linked list
func (ll *SingleLinkedNode) Insert(value interface{}, pos int) error {
	if ll.length < pos {
		return fmt.Errorf("pos %d it out of the linked list %d", pos, ll.length)
	} else if ll.Size() == pos {
		ll.Append(value)
	} else if pos == 0 {
		ll.Prepend(value)
	} else {
		node := &utils.Node{
			Value: value,
			Next:  nil,
		}

		head := ll.head
		var preNode *utils.Node
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
func (ll *SingleLinkedNode) RemoveAt(pos int) {
	if ll.length <= pos {
		fmt.Printf("pos %d it out of the linked list %d\n", pos, ll.length)
		return
	}

	head := ll.head
	var preNode *utils.Node
	for i := 0; i < ll.length; i++ {
		if i == pos {
			// 删除的位置是尾节点是，前置节点指为 nil 即可
			if head.Next == nil {
				ll.tail = preNode
				// 只有一个节点时
				if preNode == nil {
					ll.head = nil
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
func (ll *SingleLinkedNode) Delete(value interface{}) {
	head := ll.head
	var preNode *utils.Node
	for pos := 0; pos < ll.length; pos++ {
		if head.Value == value {
			// 尾节点时特殊处理
			if head.Next == nil {
				if preNode != nil {
					preNode.Next = head.Next
				} else {
					// 只有一个节点时
					ll.head = nil
				}
				ll.tail = preNode
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
func (ll *SingleLinkedNode) IndexOf(value interface{}) int {

	head := ll.head
	for pos := 0; pos < ll.length; pos++ {
		if head.Value == value {
			return pos
		}
		head = head.Next
	}
	return -1
}

// IsEmpty judge the linked list is empty
func (ll *SingleLinkedNode) IsEmpty() bool {
	if ll == nil || ll.head == nil {
		return true
	}
	return false
}

// Size get the linked list length
func (ll *SingleLinkedNode) Size() int {
	if ll.head == nil {
		return 0
	}
	return ll.length
}

// Display a string representation of the list
func (ll *SingleLinkedNode) Display() {
	if ll.head == nil {
		return
	}
	fmt.Printf("Display Single Linked List: ")

	head := ll.head
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
func (ll *SingleLinkedNode) String() string {
	if ll.head == nil {
		return "This is a empty single linked list."
	}
	var output strings.Builder
	output.WriteString("Single Linked List: ")

	head := ll.head
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
func (ll *SingleLinkedNode) Reset() {
	ll.head = nil
	ll.tail = nil
	ll.length = 0
}

// func (ll *SingleLinkedNode) Search(value interface{}) *SingleLinkedNode {
// 	for node != nil && node.Value != value {
// 		node = node.Next
// 	}
// 	return node
// }

// // HasCycle 快慢指针法
// func (ll *SingleLinkedNode) HasCycle() bool {
// 	if node == nil || node.Next == nil {
// 		return false
// 	}
// 	fastNode := node.Next
// 	for fastNode != nil && fastNode.Next != nil {
// 		node = node.Next
// 		// 注意这儿是 fastNode.Next.Next
// 		fastNode = fastNode.Next.Next

// 		// 后判断因为三个节点才能成环
// 		if fastNode == node {
// 			return true
// 		}
// 	}
// 	return false
// }

// // HasCycleWithHash 哈希法
// func (ll *SingleLinkedNode) HasCycleWithHash() bool {
// 	cache := make(map[*SingleLinkedNode]struct{})
// 	for node.Next != nil {
// 		if _, ok := cache[node]; ok {
// 			return true
// 		}
// 		cache[node] = struct{}{}
// 		node = node.Next
// 	}
// 	return false
// }
