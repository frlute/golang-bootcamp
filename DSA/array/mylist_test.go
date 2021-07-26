package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newMyList(t *testing.T) {
	l := newMyArray(1)
	assert.Equal(t, 0, l.Cap())

	l2 := newMyArray(1, 10)
	assert.Equal(t, 10, l2.Cap())
	assert.Equal(t, 0, l2.Len())
}

func Test_myList_lifecycle(t *testing.T) {
	as := assert.New(t)

	l := newMyArray(1, 3)
	l.Append(1)

	as.Equal(l.Len(), 1)
	l.Insert(1, 2)
	as.Equal(1, l.Len())

	as.Equal(0, l.GetIndex(1))
	l.Insert(0, 0)
	as.Equal(2, l.Len())

	l.Display()
	l.Append(2)
	l.Display()
	l.Append(3)

	as.Equal(4, l.Len())
	// append 动态扩容时条件加了 =，会提前一位动态扩容
	as.Equal((2+1)*2, l.Cap())

	l.Remove(2)
	as.Equal(3, l.Len())

	l.Display()
}
