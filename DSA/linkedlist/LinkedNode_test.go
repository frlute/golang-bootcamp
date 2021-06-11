package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HasCycle(t *testing.T) {
	l := NewLinkedNode(1)

	l2 := l.InsertAfter(l, 2)
	l3 := l2.InsertAfter(l2, 3)

	// Create a loop for testing
	l3.Next = l

	assert.Equal(t, true, l.HasCycle())
}
