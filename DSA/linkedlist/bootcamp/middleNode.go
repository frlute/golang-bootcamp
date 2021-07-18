package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

/*
题目： [876] Middle of the Linked List
要求：If there are two middle nodes, return the second middle node.
解题思路：
	1. 快慢指针，快指针跑到尾节点时，慢指针刚好到中间节点, 时间复杂度 O(n), 空间复杂度 O(1)
	2. 遍历法，遍历整个链表加入到数组中，然后取 length/2 时间复杂度 O(n), 空间复杂度 O(n)
	3.方法二的改进，遍历两次，第一次获取长度，第二次到中间节点时返回节点， O(n), 空间复杂度 O(1)

解题也可以参考： https://javarevisited.blogspot.com/2012/12/how-to-find-middle-element-of-linked-list-one-pass.html?source=post_page---------------------------&ref=hackernoon.com#axzz70sjwWbKO, 这里面提供了不少类似问题的链接，对比式学习更容易掌握类似问题的解法。
*/

func middleNode(head *ll.LinkedListNode) *ll.LinkedListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func middleNodeV1(head *ll.LinkedListNode) *ll.LinkedListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	pos := 0
	cur = head
	for pos < length/2 {
		pos++
		cur = cur.Next
	}

	return cur
}
