package bootcamp

import (
	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

/*
leetcode [206] Reverse Linked List

第一次考虑时没思路，看了题解，思路如下

方法一： 迭代(双指针)
Assume that we have linked list 1 → 2 → 3 → Ø, we would like to change it to Ø ← 1 ← 2 ← 3.

While you are traversing the list, change the current node's next pointer to point to its previous element. Since a node does not have reference to its previous node, you must store its previous element beforehand. You also need another pointer to store the next node before changing the reference. Do not forget to return the new head reference at the end!

时间复杂度O(n), 空间 O(1)
*/
func reverse(head *ll.LinkedListNode) *ll.LinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ll.LinkedListNode
	for head != nil {
		// 暂存下一个节点的信息
		next := head.Next
		// 反转当前链表指向
		head.Next = prev
		// 记录前一个节点
		prev = head
		// 保证迭代正常进行
		head = next
	}

	return prev
}

/*
方法二：递归
The recursive version is slightly trickier and the key is to work backwards. Assume that the rest of the list had already been reversed, now how do I reverse the front part? Let's assume the list is: n1 → … → nk-1 → nk → nk+1 → … → nm → Ø

Assume from node k+1 to nm had been reversed and you are at node nk.
n1 → … → nk-1 → nk → nk+1 ← … ← nm
We want nk+1’s next node to point to nk.

详见 https://leetcode.com/problems/reverse-linked-list/solution/
*/
func reverseBasedRecursive(head *ll.LinkedListNode) *ll.LinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseBasedRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return p
}
