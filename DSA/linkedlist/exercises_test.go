package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	as := assert.New(t)

	ll.Append("one")
	ll.Append("two")
	ll.Append("three")

	// Create a loop for testing
	tailNode := ll.Tail
	tailNode.Next = ll.Head

	// TODO 暂未造多余的数据
	as.Equal(true, hasCycle(ll.Head))
	as.Equal(true, hasCycleBasedHash(ll.Head))
}

func Test_removeNthFromEnd(t *testing.T) {
	as := assert.New(t)

	cases := []struct {
		input  []int
		pos    int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			pos:    2,
			output: []int{1, 2, 3, 5},
		},
		{
			input:  []int{1},
			pos:    1,
			output: []int{},
		},
		{
			input:  []int{1, 2},
			pos:    1,
			output: []int{1},
		},
	}

	for _, args := range cases {
		input := numberArrayToLinkedlist(args.input)

		res := removeNthFromEnd(input, args.pos)
		list := &SingleLinkedList{Head: res}

		expectNode := numberArrayToLinkedlist(args.output)
		expectList := &SingleLinkedList{Head: expectNode}

		as.Equal(expectList.String(), list.String())
	}
}

func Test_reverse(t *testing.T) {
	as := assert.New(t)

	cases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{5, 4, 3, 2, 1},
		},
		{
			input:  []int{1, 2},
			output: []int{2, 1},
		},
		{
			input:  []int{},
			output: []int{},
		},
	}

	for _, args := range cases {
		name := fmt.Sprintf("反转链表%+v", args.input)
		t.Run(name, func(t *testing.T) {
			input := numberArrayToLinkedlist(args.input)

			res := reverse(input)
			list := &SingleLinkedList{Head: res}

			expectNode := numberArrayToLinkedlist(args.output)
			expectList := &SingleLinkedList{Head: expectNode}

			as.Equal(expectList.String(), list.String())
		})
	}

	for _, args := range cases {
		name := fmt.Sprintf("递归反转链表%+v", args.input)
		t.Run(name, func(t *testing.T) {
			input := numberArrayToLinkedlist(args.input)

			res := reverseBasedRecursive(input)
			list := &SingleLinkedList{Head: res}

			expectNode := numberArrayToLinkedlist(args.output)
			expectList := &SingleLinkedList{Head: expectNode}

			as.Equal(expectList.String(), list.String())
		})
	}
}

func Test_oddEvenList(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 3, 5, 2, 4},
		},
		{
			input:  []int{2, 1, 3, 5, 6, 4, 7},
			output: []int{2, 3, 6, 7, 1, 5, 4},
		},
	}

	for _, args := range cases {
		name := fmt.Sprintf("oddEven链表%+v", args.input)
		t.Run(name, func(t *testing.T) {
			input := numberArrayToLinkedlist(args.input)

			res := oddEvenList(input)
			list := &SingleLinkedList{Head: res}

			expectNode := numberArrayToLinkedlist(args.output)
			expectList := &SingleLinkedList{Head: expectNode}

			assert.Equal(t, expectList.String(), list.String())
		})
	}
}

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		output []int
	}{
		{
			input1: []int{2, 4, 3},
			input2: []int{5, 6, 4},
			output: []int{7, 0, 8},
		},
		{
			input1: []int{0},
			input2: []int{0},
			output: []int{0},
		},
		{
			input1: []int{9, 9, 9, 9, 9, 9, 9},
			input2: []int{9, 9, 9, 9},
			output: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, test := range tests {
		input1 := numberArrayToLinkedlist(test.input1)
		input2 := numberArrayToLinkedlist(test.input2)

		res := addTwoNumbers(input1, input2)

		list := &SingleLinkedList{Head: res}

		expectNode := numberArrayToLinkedlist(test.output)
		expectList := &SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())

	}
}

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4},
			output: []int{2, 1, 4, 3},
		},
		{
			input:  []int{},
			output: []int{},
		},
		{
			input:  []int{1},
			output: []int{1},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			output: []int{2, 1, 4, 3, 5},
		},
	}

	for _, test := range tests {
		input := numberArrayToLinkedlist(test.input)

		res := swapPairsIteration(input)
		list := &SingleLinkedList{Head: res}

		expectNode := numberArrayToLinkedlist(test.output)
		expectList := &SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())
	}

	// 递归法
	for _, test := range tests {
		input := numberArrayToLinkedlist(test.input)

		res := swapPairs(input)
		list := &SingleLinkedList{Head: res}

		expectNode := numberArrayToLinkedlist(test.output)
		expectList := &SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())
	}
}

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 1, 2},
			output: []int{1, 2},
		},
		{
			input:  []int{1, 1, 2, 3, 3},
			output: []int{1, 2, 3},
		},
		{
			input:  []int{1, 1, 1},
			output: []int{1},
		},
	}

	for _, test := range tests {
		input := numberArrayToLinkedlist(test.input)

		res := deleteDuplicates(input)
		list := &SingleLinkedList{Head: res}

		expectNode := numberArrayToLinkedlist(test.output)
		expectList := &SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())
	}

	for _, test := range tests {
		input := numberArrayToLinkedlist(test.input)

		res := deleteDuplicatesV1(input)
		list := &SingleLinkedList{Head: res}

		expectNode := numberArrayToLinkedlist(test.output)
		expectList := &SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())
	}
}
