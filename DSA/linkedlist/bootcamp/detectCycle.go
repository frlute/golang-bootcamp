package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

/*
* [142] Linked List Cycle II
题目：判断链表是否成环及成环点

解题思路：
	1. 哈希法
	2. 快慢指针
	3.

其他题解：https://javarevisited.blogspot.sg/2013/05/find-if-linked-list-contains-loops-cycle-cyclic-circular-check.html?source=post_page---------------------------&ref=hackernoon.com
*/

func detectCycle(head *ll.LinkedListNode) *ll.LinkedListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}

	return nil
}
