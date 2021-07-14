package bootcamp

import (
	"fmt"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

/*
题目: Swap Nodes in Pairs
地址 https://leetcode.com/problems/swap-nodes-in-pairs/
*/

// 递归法, 时间复杂度 O(n), 空间复杂度 O(n)
func swapPairs(head *ll.LinkedListNode) *ll.LinkedListNode {
	// 递归终止条件，没有成对的节点
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}

// 迭代法 时间复杂度 O(n), 空间复杂度 O(1)
// TODO 没明白。。。。
func swapPairsIteration(head *ll.LinkedListNode) *ll.LinkedListNode {
	dummyHead := &ll.LinkedListNode{nil, head}
	temp := dummyHead
	// 注意这儿的条件，是从 .next.next 作为 dump 的起始节点
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next

		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		temp = node1
	}

	test := &ll.SingleLinkedList{Head: dummyHead}
	fmt.Println("test: ")
	test.Display()

	return dummyHead.Next
}
