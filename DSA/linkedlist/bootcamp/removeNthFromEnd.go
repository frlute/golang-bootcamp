package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

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
func removeNthFromEnd(head *ll.LinkedListNode, pos int) *ll.LinkedListNode {
	fastNode, slowNode := head, head
	// 首先移到从尾节点到 n 对应的节点
	for ; pos > 0; pos-- {
		fastNode = fastNode.Next
	}

	// 因 n 大于等于链表长度，此时相当于 n = len(ll), 即删除头结点
	if fastNode == nil {
		return head.Next
	}

	for fastNode.Next != nil && fastNode != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}

	slowNode.Next = slowNode.Next.Next
	return head
}
