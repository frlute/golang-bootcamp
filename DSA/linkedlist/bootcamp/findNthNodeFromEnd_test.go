package bootcamp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

func Test_findNthNodeFromEnd(t *testing.T) {
	as := assert.New(t)

	cases := []struct {
		input  []int
		pos    int
		output []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			pos:    2,
			output: []int{4, 5},
		},
		{
			input:  []int{1},
			pos:    1,
			output: []int{1},
		},
		{
			input:  []int{1, 2},
			pos:    1,
			output: []int{2},
		},
	}

	for _, args := range cases {
		input := ll.NumberArrayToLinkedlist(args.input)

		res := findNthNodeFromEnd(input, args.pos)
		list := &ll.SingleLinkedList{Head: res}
		fmt.Println("execute result: ")
		list.Display()

		expectNode := ll.NumberArrayToLinkedlist(args.output)
		expectList := &ll.SingleLinkedList{Head: expectNode}
		fmt.Println("expect result: ")
		expectList.Display()

		as.Equal(expectList.String(), list.String())
	}
}
