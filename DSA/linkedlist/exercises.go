package linkedlist

import "fmt"

// 此文件主要用来做练习题
/*
[141] Linked List Cycle
*/
func hasCycle(head *LinkedListNode) bool {
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

func hasCycleBasedHash(head *LinkedListNode) bool {
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
func removeNthFromEnd(head *LinkedListNode, pos int) *LinkedListNode {
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

/*
leetcode [206] Reverse Linked List

第一次考虑时没思路，看了题解，思路如下

方法一： 迭代(双指针)
Assume that we have linked list 1 → 2 → 3 → Ø, we would like to change it to Ø ← 1 ← 2 ← 3.

While you are traversing the list, change the current node's next pointer to point to its previous element. Since a node does not have reference to its previous node, you must store its previous element beforehand. You also need another pointer to store the next node before changing the reference. Do not forget to return the new head reference at the end!

时间复杂度O(n), 空间 O(1)
*/
func reverse(head *LinkedListNode) *LinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *LinkedListNode
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
func reverseBasedRecursive(head *LinkedListNode) *LinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseBasedRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return p
}

/*
leetcode [328] Odd Even Linked List
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

详见： https://leetcode.com/problems/odd-even-linked-list/

解题思路：
	方法一: 分离节点后合并
*/
func oddEvenList(head *LinkedListNode) *LinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddNode := head
	evenNode := head.Next

	even := evenNode
	for even != nil && even.Next != nil {
		// 奇数更新
		// 1.Next = 2.Next = 3
		oddNode.Next = even.Next
		// 1-> 3
		oddNode = oddNode.Next

		// 2.Next = 3.Next = 4
		even.Next = oddNode.Next
		// 2 -> 4
		even = even.Next
	}

	// merge
	oddNode.Next = evenNode
	return head
}

/*
题目: 两数之和，input 以逆序方式提供
地址：[2] https://leetcode.com/problems/add-two-numbers/
时间复杂度 O(max(m ,n))
*/
func addTwoNumbers(l1 *LinkedListNode, l2 *LinkedListNode) *LinkedListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	// 利用了哨兵节点
	dummy := &LinkedListNode{0, nil}
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

		dummy.Next = &LinkedListNode{sum % 10, nil}
		// 大于 10 时进一位
		sum = sum / 10
		dummy = dummy.Next
	}

	if sum > 0 {
		dummy.Next = &LinkedListNode{sum, nil}
	}

	return head.Next
}

/*
题目: Swap Nodes in Pairs
地址 https://leetcode.com/problems/swap-nodes-in-pairs/
*/

// 递归法, 时间复杂度 O(n), 空间复杂度 O(n)
func swapPairs(head *LinkedListNode) *LinkedListNode {
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
func swapPairsIteration(head *LinkedListNode) *LinkedListNode {
	dummyHead := &LinkedListNode{nil, head}
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

	test := &SingleLinkedList{Head: dummyHead}
	fmt.Println("test: ")
	test.Display()

	return dummyHead.Next
}
