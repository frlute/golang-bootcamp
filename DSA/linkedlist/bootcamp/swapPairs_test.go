package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

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
		input := ll.NumberArrayToLinkedlist(test.input)

		res := swapPairsIteration(input)
		list := &ll.SingleLinkedList{Head: res}

		expectNode := ll.NumberArrayToLinkedlist(test.output)
		expectList := &ll.SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())
	}

	// 递归法
	for _, test := range tests {
		input := ll.NumberArrayToLinkedlist(test.input)

		res := swapPairs(input)
		list := &ll.SingleLinkedList{Head: res}

		expectNode := ll.NumberArrayToLinkedlist(test.output)
		expectList := &ll.SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())
	}
}
