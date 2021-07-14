package bootcamp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

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
			input := ll.NumberArrayToLinkedlist(args.input)

			res := reverse(input)
			list := &ll.SingleLinkedList{Head: res}

			expectNode := ll.NumberArrayToLinkedlist(args.output)
			expectList := &ll.SingleLinkedList{Head: expectNode}

			as.Equal(expectList.String(), list.String())
		})
	}

	for _, args := range cases {
		name := fmt.Sprintf("递归反转链表%+v", args.input)
		t.Run(name, func(t *testing.T) {
			input := ll.NumberArrayToLinkedlist(args.input)

			res := reverseBasedRecursive(input)
			list := &ll.SingleLinkedList{Head: res}

			expectNode := ll.NumberArrayToLinkedlist(args.output)
			expectList := &ll.SingleLinkedList{Head: expectNode}

			as.Equal(expectList.String(), list.String())
		})
	}
}
