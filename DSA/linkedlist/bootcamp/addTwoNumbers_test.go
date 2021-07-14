package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

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
		input1 := ll.NumberArrayToLinkedlist(test.input1)
		input2 := ll.NumberArrayToLinkedlist(test.input2)

		res := addTwoNumbers(input1, input2)

		list := &ll.SingleLinkedList{Head: res}

		expectNode := ll.NumberArrayToLinkedlist(test.output)
		expectList := &ll.SingleLinkedList{Head: expectNode}

		assert.Equal(t, expectList.String(), list.String())

	}
}
