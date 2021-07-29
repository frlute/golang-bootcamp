package bootcamp

import ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"

/*
题目： 发现倒数第 N 个节点
*/

func findNthNodeFromEnd(head *ll.LinkedListNode, pos int) *ll.LinkedListNode {

	fastNode, slowNode := head, head
	start := 1

	for fastNode.Next != nil {
		fastNode = fastNode.Next
		start++
		if start > pos {
			slowNode = slowNode.Next
		}
	}

	return slowNode
}
