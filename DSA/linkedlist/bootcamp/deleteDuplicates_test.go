package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

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
		input := ll.NumberArrayToLinkedlist(test.input)

		res := deleteDuplicates(input)
		list := &ll.SingleLinkedList{Head: res}

		expectNode := ll.NumberArrayToLinkedlist(test.output)
		expectList := &ll.SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())
	}

	for _, test := range tests {
		input := ll.NumberArrayToLinkedlist(test.input)

		res := deleteDuplicatesV1(input)
		list := &ll.SingleLinkedList{Head: res}

		expectNode := ll.NumberArrayToLinkedlist(test.output)
		expectList := &ll.SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())
	}
}
