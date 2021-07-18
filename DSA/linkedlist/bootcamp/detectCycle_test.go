package bootcamp

import (
	"testing"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
	"github.com/stretchr/testify/assert"
)

func Test_detectCycle(t *testing.T) {
	tests := []struct {
		input  []int
		pos    int // 成环点
		output []int
	}{
		{
			input:  []int{3, 2, 0, -4},
			pos:    1,
			output: []int{2, 0, -4},
		},
		{
			input:  []int{1, 2},
			pos:    0,
			output: []int{1, 2},
		},
		{
			input:  []int{1},
			pos:    -1,
			output: nil,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			// makeTestData(test.input, test.pos)
			head := makeTestData(test.input, test.pos)
			res := detectCycle(head)

			list := &ll.SingleLinkedList{Head: res}

			expectNode := ll.NumberArrayToLinkedlist(test.output)
			expectList := &ll.SingleLinkedList{Head: expectNode}

			assert.Equal(t, expectList.String(), list.String())
		})
	}
}

func makeTestData(input []int, cyclePos int) *ll.LinkedListNode {
	var inputList ll.SingleLinkedList
	for _, item := range input {
		inputList.Append(item)
	}

	cycleNode := inputList.Index(cyclePos)
	inputList.AppendIntoTail(cycleNode)
	inputList.Display()

	return inputList.Head
}
