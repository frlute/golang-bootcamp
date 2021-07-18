package bootcamp

import (
	"fmt"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

type inputFunc func(*ll.LinkedListNode) *ll.LinkedListNode

type executeResult struct {
	result ll.LinkedList
	expect ll.LinkedList
}

func runIntSingleLinkedListFunc(input, output []int, fn inputFunc) executeResult {
	inputNode := ll.NumberArrayToLinkedlist(input)

	result := fn(inputNode)
	list := &ll.SingleLinkedList{Head: result}
	fmt.Println("result data: ")
	list.Display()

	expectNode := ll.NumberArrayToLinkedlist(output)
	expectList := &ll.SingleLinkedList{Head: expectNode}

	fmt.Println("expect output data: ")
	expectList.Display()

	return executeResult{
		list,
		expectList,
	}
}
