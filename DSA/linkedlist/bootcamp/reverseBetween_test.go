package bootcamp

import (
	"fmt"
	"testing"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
	"github.com/stretchr/testify/assert"
)

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		input  []int
		left   int
		right  int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6},
			left:   3,
			right:  5,
			output: []int{1, 2, 5, 4, 3, 6},
		},
		{
			input:  []int{5},
			left:   1,
			right:  1,
			output: []int{5},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			inputNode := ll.NumberArrayToLinkedlist(test.input)

			res := reverseBetween(inputNode, test.left, test.right)

			list := &ll.SingleLinkedList{Head: res}
			fmt.Println("result data: ")
			list.Display()

			expectNode := ll.NumberArrayToLinkedlist(test.output)
			expectList := &ll.SingleLinkedList{Head: expectNode}

			fmt.Println("expect output data: ")
			expectList.Display()

			assert.Equal(t, list.String(), expectList.String())
		})
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			inputNode := ll.NumberArrayToLinkedlist(test.input)

			res1 := reverseBetweenIteration(inputNode, test.left, test.right)

			list1 := &ll.SingleLinkedList{Head: res1}
			fmt.Println("list1 data: ")
			list.Display()

			expectNode := ll.NumberArrayToLinkedlist(test.output)
			expectList := &ll.SingleLinkedList{Head: expectNode}

			fmt.Println("expect output data: ")
			expectList.Display()

			assert.Equal(t, list1.String(), expectList.String())
		})
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			inputNode := ll.NumberArrayToLinkedlist(test.input)

			res2 := reverseBetweenV2(inputNode, test.left, test.right)

			list2 := &ll.SingleLinkedList{Head: res2}
			fmt.Println("list2 data: ")
			list.Display()

			expectNode := ll.NumberArrayToLinkedlist(test.output)
			expectList := &ll.SingleLinkedList{Head: expectNode}

			fmt.Println("expect output data: ")
			expectList.Display()

			assert.Equal(t, list2.String(), expectList.String())
		})
	}
}
