package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JosephusProblem(t *testing.T) {
	soulution := newJosephusSolutaion(12)
	soulution.start(3)
	assert.Equal(t, "display circular Linked List: 1\n", soulution.cll.string())

	soulution2 := newJosephusSolutaion(8)
	soulution2.start(3)
	assert.Equal(t, "display circular Linked List: 6\n", soulution2.cll.string())
}
