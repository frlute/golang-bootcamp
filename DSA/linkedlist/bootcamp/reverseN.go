package bootcamp

import ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"

/*
问题： 将链表的前 n 个节点反转（n <= 链表长度）
思路：
	* 首先找到第 n 个节点

方法：
	* 迭代法
	* 递归法
*/
var postNode *ll.LinkedListNode

func reverseN(head *ll.LinkedListNode, n int) *ll.LinkedListNode {
	// 如果链表只有一个元素直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 递归终止条件
	if n == 1 {
		// 记录第 n + 1 个节点
		postNode = head.Next
		return head
	}

	// 以 lastNode 为七点，反转前 n-1个节点
	lastNode := reverseN(head.Next, n-1)
	head.Next.Next = head
	// 连接反转之后的 tail 节点与后面未反转的节点
	head.Next = postNode

	return lastNode
}
