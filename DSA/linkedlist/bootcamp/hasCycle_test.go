package bootcamp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ll "github.com/frlute/golang-bootcamp/DSA/linkedlist"
)

var list ll.SingleLinkedList

func Test_hasCycle(t *testing.T) {
	as := assert.New(t)

	list.Append("one")
	list.Append("two")
	list.Append("three")

	// Create a loop for testing
	tailNode := list.Tail
	tailNode.Next = list.Head

	// TODO 暂未造多余的数据
	as.Equal(true, hasCycle(list.Head))
	as.Equal(true, hasCycleBasedHash(list.Head))
}
