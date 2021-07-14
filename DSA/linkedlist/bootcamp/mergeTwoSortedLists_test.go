package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

func Test_mergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		output []int
	}{
		{
			input1: []int{1, 2, 4},
			input2: []int{1, 3, 4},
			output: []int{1, 1, 2, 3, 4, 4},
		},
		{
			input1: []int{},
			input2: []int{},
			output: []int{},
		},
		{
			input1: []int{},
			input2: []int{0},
			output: []int{0},
		},
	}

	for _, test := range tests {
		input1 := ll.NumberArrayToLinkedlist(test.input1)
		input2 := ll.NumberArrayToLinkedlist(test.input2)

		res := mergeTwoSortedLists(input1, input2)

		list := &ll.SingleLinkedList{Head: res}

		expectNode := ll.NumberArrayToLinkedlist(test.output)
		expectList := &ll.SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())
	}
}
