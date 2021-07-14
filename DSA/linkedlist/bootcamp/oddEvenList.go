package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

/*
leetcode [328] Odd Even Linked List
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

详见： https://leetcode.com/problems/odd-even-linked-list/

解题思路：
	方法一: 分离节点后合并
*/
func oddEvenList(head *ll.LinkedListNode) *ll.LinkedListNode {
	if head == nil || head.Next == nil {
		return head
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
	return head
}
