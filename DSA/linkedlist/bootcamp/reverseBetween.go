package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

/*

题目： 反转指定区间的链表 [92] Reverse Linked List II
思路：
	* 递归法
		1. 如果 m == 1，就相当于反转链表开头的 n 个元素嘛
	    2. 如果 m != 1 怎么办？如果我们把 head 的索引视为 1，那么我们是想从第 m 个元素开始反转对吧；如果把 head.next 的索引视为 1 呢？那么相对于 head.next，反转的区间应该是从第 m - 1 个元素开始的；那么对于 head.next.next 呢……


简单明了的解释见： https://zhuanlan.zhihu.com/p/86745433
*/

func reverseBetween(head *ll.LinkedListNode, left int, right int) *ll.LinkedListNode {
	// left=1时，表示反转链表的第n个节点
	if left == 1 {
		return reverseN(head, right)
	}
	// 不需要反转则一直递归下去
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 迭代法, 解法可参考 https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/fan-zhuan-lian-biao-ii-by-leetcode-solut-teyq/
func reverseBetweenIteration(head *ll.LinkedListNode, left, right int) *ll.LinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 哨兵节点
	dummy := &ll.LinkedListNode{}
	dummy.Next = head
	//指向新头结点
	pre := dummy

	// 从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206 题，反转链表的子区间
	reverse(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr

	return dummy.Next
}

func reverseBetweenV2(head *ll.LinkedListNode, left, right int) *ll.LinkedListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ll.LinkedListNode{}
	dummyNode.Next = head

	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
