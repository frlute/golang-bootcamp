package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

// 此文件主要用来做练习题
/*
[141] Linked List Cycle
*/
// TODO 感觉实现不对，待验证, 也不能说是实现不对，只是成环点确实存在问题，这种实现只是能知道成环了，但是具体的成环点存在问题
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

	cached := make(map[*ll.LinkedListNode]struct{})
	for head != nil {
		if _, ok := cached[head]; ok {
			return true
		}
		cached[head] = struct{}{}
		head = head.Next
	}
	return false
}
