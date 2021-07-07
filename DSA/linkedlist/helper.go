package linkedlist

func numberArrayToLinkedlist(items []int) *LinkedListNode {
	var list SingleLinkedList
	for _, item := range items {
		list.Append(item)
	}

	return list.Head
}
