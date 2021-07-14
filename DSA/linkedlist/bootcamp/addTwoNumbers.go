package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

/*
题目: 两数之和，input 以逆序方式提供
地址：[2] https://leetcode.com/problems/add-two-numbers/
时间复杂度 O(max(m ,n))
*/
func addTwoNumbers(l1 *ll.LinkedListNode, l2 *ll.LinkedListNode) *ll.LinkedListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	// 利用了哨兵节点
	dummy := &ll.LinkedListNode{0, nil}
	// 返回链表头结点
	head := dummy

	// 进位符
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Value.(int)
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Value.(int)
			l2 = l2.Next
		}

		dummy.Next = &ll.LinkedListNode{sum % 10, nil}
		// 大于 10 时进一位
		sum = sum / 10
		dummy = dummy.Next
	}

	if sum > 0 {
		dummy.Next = &ll.LinkedListNode{sum, nil}
	}

	return head.Next
}
