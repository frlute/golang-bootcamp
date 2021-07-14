package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

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
		input := ll.NumberArrayToLinkedlist(args.input)

		res := removeNthFromEnd(input, args.pos)
		list := &ll.SingleLinkedList{Head: res}

		expectNode := ll.NumberArrayToLinkedlist(args.output)
		expectList := &ll.SingleLinkedList{Head: expectNode}

		as.Equal(expectList.String(), list.String())
	}
}
