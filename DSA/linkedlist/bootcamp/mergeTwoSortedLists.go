package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

/*
 * [21] Merge Two Sorted Lists
 */
// 迭代法，时间复杂度 O(m+n), 空间复杂度 O(1)
func mergeTwoSortedLists(l1, l2 *ll.LinkedListNode) *ll.LinkedListNode {
	// 哨兵节点
	result := &ll.LinkedListNode{}

	changeNode := result
	for l1 != nil && l2 != nil {
		if l1.Value.(int) > l2.Value.(int) {
			changeNode.Next = l2
			l2 = l2.Next
		} else {
			changeNode.Next = l1
			l1 = l1.Next
		}
		changeNode = changeNode.Next
	}

	// 如果 l1 与 l2 长度不一致，则增加剩余的部分
	if l1 != nil {
		changeNode.Next = l1
	}
	if l2 != nil {
		changeNode.Next = l2
	}

	return result.Next
}

// TODO 待整理思路 递归法
func mergeTwoSortedListsV1(l1, l2 *ll.LinkedListNode) *ll.LinkedListNode {
	if l1 == nil {
		return l1
	} else if l2 == nil {
		return l2
	}

	if l1.Value.(int) < l2.Value.(int) {
		l1.Next = mergeTwoSortedListsV1(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoSortedListsV1(l2.Next, l1)
	return l2
}
