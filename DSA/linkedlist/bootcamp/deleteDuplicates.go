package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

/*
[83] Remove Duplicates from Sorted List
方法： 哈希
算法分析: 时间复杂度 O(n), 空间复杂度 O(n)
*/
func deleteDuplicates(head *ll.LinkedListNode) *ll.LinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	result := prev

	cached := make(map[interface{}]struct{})
	for head != nil {
		if _, ok := cached[head.Value]; ok {
			prev.Next = head.Next
		} else {
			cached[head.Value] = struct{}{}
			prev = head
		}
		head = head.Next
	}

	return result
}

// 迭代法: 时间复杂度 O(n), 空间复杂度 O(1)
func deleteDuplicatesV1(head *ll.LinkedListNode) *ll.LinkedListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Value == cur.Next.Value {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
