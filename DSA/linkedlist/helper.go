package linkedlist

// NumberArrayToLinkedlist _
func NumberArrayToLinkedlist(items []int) *LinkedListNode {
	var list SingleLinkedList
	for _, item := range items {
		list.Append(item)
	}

	return list.Head
}
