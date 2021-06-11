package linkedlist

// TODO 优化，把 linkedlist 抽象成 interface，此时用 singleLinkedList 表示，这样所有的链表方法就统一了
// LinkedNode Definition for singly-linked list.
type LinkedNode struct {
	Value interface{}
	Next  *LinkedNode
}

func NewLinkedNode(value interface{}) *LinkedNode {
	return &LinkedNode{
		Value: value,
		Next:  nil,
	}
}

// Insert a new node after a specified node
func (node *LinkedNode) InsertAfter(preNode *LinkedNode, value interface{}) *LinkedNode {
	newNode := NewLinkedNode(value)
	newNode.Next = node.Next
	node.Next = newNode

	return newNode
}

// func(node *LinkedNode)Push(value interface{}) {
// 	newNode := NewLinkedNode(value)

// }

func (node *LinkedNode) Search(value interface{}) *LinkedNode {
	for node != nil && node.Value != value {
		node = node.Next
	}
	return node
}

func (node *LinkedNode) Delete(value interface{}) {
	node.Next = node.Next.Next
}

// HasCycle 快慢指针法
func (node *LinkedNode) HasCycle() bool {
	if node == nil || node.Next == nil {
		return false
	}
	fastNode := node.Next
	for fastNode != nil && fastNode.Next != nil {
		node = node.Next
		// 注意这儿是 fastNode.Next.Next
		fastNode = fastNode.Next.Next

		// 后判断因为三个节点才能成环
		if fastNode == node {
			return true
		}
	}
	return false
}

// HasCycleWithHash 哈希法
func (node *LinkedNode) HasCycleWithHash() bool {
	cache := make(map[*LinkedNode]struct{})
	for node.Next != nil {
		if _, ok := cache[node]; ok {
			return true
		}
		cache[node] = struct{}{}
		node = node.Next
	}
	return false
}
