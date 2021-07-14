package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

// 此文件主要用来做练习题
/*
[141] Linked List Cycle
*/
func hasCycle(head *ll.LinkedListNode) bool {
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

func hasCycleBasedHash(head *ll.LinkedListNode) bool {
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
