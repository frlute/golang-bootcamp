package bootcamp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

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
			input := ll.NumberArrayToLinkedlist(args.input)

			res := oddEvenList(input)
			list := &ll.SingleLinkedList{Head: res}

			expectNode := ll.NumberArrayToLinkedlist(args.output)
			expectList := &ll.SingleLinkedList{Head: expectNode}

			assert.Equal(t, expectList.String(), list.String())
		})
	}
}
