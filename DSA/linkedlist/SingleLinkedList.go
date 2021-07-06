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
	// 记录 tail 节点方便部分操作
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

/*------------------以下为练习对应的实现------------*/

// 快慢指针法
func (ll *SingleLinkedList) hasCycle() bool {
	head := ll.Head
	if head == nil || head.Next == nil {
		return false
	}

	fastNode := head.Next
	for fastNode != nil && fastNode.Next != nil {
		head = head.Next
		// 注意这儿是 fastNode.Next.Next, 三个点成环
		fastNode = fastNode.Next.Next

		if fastNode == head {
			return true
		}
	}

	return false
}

// 哈希法
func (ll *SingleLinkedList) hasCycleBasedHash() bool {
	head := ll.Head
	if head == nil || head.Next == nil {
		return false
	}

	cached := make(map[*LinkedListNode]struct{}, ll.length)
	for head.Next != nil {
		if _, ok := cached[head]; ok {
			return true
		}
		cached[head] = struct{}{}
		head = head.Next

	}
	return false
}

/*
leetcode [19] Remove Nth Node From End of List

限制条件:
	* The number of nodes in the list is sz.
	* 1 <= sz <= 30
	* 0 <= Node.val <= 100
	* 1 <= n <= sz

解法：
	- 双指针法
	- 先确定好节点位置，然后遍历当节点 == nil 时说明已经确定的求解头结点的位置
*/
func (ll *SingleLinkedList) removeNthFromEnd(pos int) {
	fastNode, slowNode := ll.Head, ll.Head
	// 首先移到从尾节点到 n 对应的节点
	for ; pos > 0; pos-- {
		fastNode = fastNode.Next
	}

	// 为了方便测试，通常实现没有此行
	ll.length--
	// 因 n 大于等于链表长度，此时相当于 n = len(ll), 即删除头结点
	if fastNode == nil {
		ll.Head = ll.Head.Next
		return
	}

	for fastNode.Next != nil && fastNode != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}

	slowNode.Next = slowNode.Next.Next
}

/*
leetcode [206] Reverse Linked List

第一次考虑时没思路，看了题解，思路如下

方法一： 迭代(双指针)
Assume that we have linked list 1 → 2 → 3 → Ø, we would like to change it to Ø ← 1 ← 2 ← 3.

While you are traversing the list, change the current node's next pointer to point to its previous element. Since a node does not have reference to its previous node, you must store its previous element beforehand. You also need another pointer to store the next node before changing the reference. Do not forget to return the new head reference at the end!

时间复杂度O(n), 空间 O(1)
*/
func (ll *SingleLinkedList) reverse() {
	head := ll.Head
	if head == nil {
		return
	}
	var prev *LinkedListNode
	for head != nil {
		// 暂存下一个节点的信息
		next := head.Next
		// 反转当前链表指向
		head.Next = prev
		// 记录前一个节点
		prev = head
		// 保证迭代正常进行
		head = next
	}
	// 尾节点变成了头结点
	ll.Head = prev
}

/*
方法二：递归
The recursive version is slightly trickier and the key is to work backwards. Assume that the rest of the list had already been reversed, now how do I reverse the front part? Let's assume the list is: n1 → … → nk-1 → nk → nk+1 → … → nm → Ø

Assume from node k+1 to nm had been reversed and you are at node nk.
n1 → … → nk-1 → nk → nk+1 ← … ← nm
We want nk+1’s next node to point to nk.

详见 https://leetcode.com/problems/reverse-linked-list/solution/
*/
func (ll *SingleLinkedList) reverseBasedRecursive(head *LinkedListNode) *LinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := ll.reverseBasedRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return p
}

/*
leetcode [328] Odd Even Linked List
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

详见： https://leetcode.com/problems/odd-even-linked-list/

解题思路：
	方法一: 分离节点后合并
*/
func (ll *SingleLinkedList) oddEvenList() {
	head := ll.Head
	if head == nil || head.Next == nil {
		return
	}

	oddNode := head
	evenNode := head.Next

	even := evenNode
	for even != nil && even.Next != nil {
		// 奇数更新
		// 1.Next = 2.Next = 3
		oddNode.Next = even.Next
		// 1-> 3
		oddNode = oddNode.Next

		// 2.Next = 3.Next = 4
		even.Next = oddNode.Next
		// 2 -> 4
		even = even.Next
	}

	// merge
	oddNode.Next = evenNode
}
