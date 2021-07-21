package bootcamp

import (
	"fmt"
	"testing"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
	"github.com/stretchr/testify/assert"
)

func Test_reverseN(t *testing.T) {
	tests := []struct {
		input  []int
		n      int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6},
			n:      3,
			output: []int{3, 2, 1, 4, 5, 6},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			inputNode := ll.NumberArrayToLinkedlist(test.input)

			res := reverseN(inputNode, test.n)

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
}
