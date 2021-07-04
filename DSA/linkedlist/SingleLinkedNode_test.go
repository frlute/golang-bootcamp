package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO ll 不能是指针类型？
var ll SingleLinkedNode

func Test_Append(t *testing.T) {
	as := assert.New(t)

	as.Equal(true, ll.IsEmpty())
	ll.Append(1)
	as.Equal(1, ll.Size())

	ll.Append(2)
	as.Equal(2, ll.Size())

	as.Equal("Single Linked List: 1->2\n", ll.String())
	as.Equal(false, ll.IsEmpty())
	as.Equal(2, ll.tail.Value)

	ll.Reset()
}

func Test_Prepend(t *testing.T) {
	as := assert.New(t)

	as.Equal(true, ll.IsEmpty())
	ll.Prepend("one")
	as.Equal(1, ll.Size())
	as.Equal(false, ll.IsEmpty())

	ll.Prepend("zero")
	as.Equal("Single Linked List: zero->one\n", ll.String())

	ll.Reset()
}

func Test_Insert(t *testing.T) {
	as := assert.New(t)

	as.Equal(true, ll.IsEmpty())
	as.Equal(0, ll.Size())
	as.Error(ll.Insert("test", 10))
	ll.Insert("one", 0)
	as.Equal(1, ll.Size())

	ll.Append("end")

	ll.Insert("two", 1)
	as.Equal(3, ll.Size())

	ll.Insert("zero", 0)
	as.Equal(4, ll.Size())

	// 验证 tail 节点是否正确
	as.Equal("end", ll.tail.Value)
	as.Nil(ll.tail.Next)

	as.Equal("Single Linked List: zero->one->two->end\n", ll.String())
	ll.Reset()
}

func Test_RemoveAt(t *testing.T) {
	as := assert.New(t)
	as.Equal(true, ll.IsEmpty())
	ll.RemoveAt(10)

	ll.Append("one")
	ll.Prepend("zero")
	ll.Append("two")
	as.Equal(3, ll.Size())

	// 删中间节点
	ll.RemoveAt(1)
	as.Equal(2, ll.Size())
	as.Equal("Single Linked List: zero->two\n", ll.String())

	// 删尾节点(考虑只有一个元素)
	ll.RemoveAt(1)
	as.Equal(1, ll.Size())
	as.Equal("Single Linked List: zero\n", ll.String())

	// 删首节点(考虑只有一个元素)
	ll.RemoveAt(0)
	as.Equal(0, ll.Size())
	as.Equal("This is a empty single linked list.", ll.String())
	as.Nil(ll.tail)
}

func Test_Delete(t *testing.T) {
	as := assert.New(t)

	ll.Append("one")
	ll.Append("two")
	ll.Append("three")
	as.Equal(3, ll.Size())

	ll.Delete("test")
	as.Equal("Single Linked List: one->two->three\n", ll.String())

	// 删中间节点
	ll.Delete("two")
	as.Equal(2, ll.Size())
	as.Equal("Single Linked List: one->three\n", ll.String())
	as.Equal("three", ll.tail.Value)

	// 删尾节点
	ll.Delete("three")
	as.Equal(1, ll.Size())
	as.Equal("Single Linked List: one\n", ll.String())
	as.Equal("one", ll.tail.Value)

	// 删头节点
	ll.Delete("one")
	as.Equal(0, ll.Size())
	as.Equal("This is a empty single linked list.", ll.String())
	as.Nil(ll.tail)

	ll.Display()
}

func Test_IndexOf(t *testing.T) {
	as := assert.New(t)

	ll.Append("one")
	ll.Append("two")
	ll.Append("three")
	as.Equal(3, ll.Size())

	as.Equal(0, ll.IndexOf("one"))
	as.Equal(2, ll.IndexOf("three"))
	as.Equal(-1, ll.IndexOf("five"))
}

// func Test_HasCycle(t *testing.T) {
// 	l := NewLinkedNode(1)

// 	l2 := l.InsertAfter(l, 2)
// 	l3 := l2.InsertAfter(l2, 3)

// 	// Create a loop for testing
// 	l3.Next = l

// 	assert.Equal(t, true, l.HasCycle())
// }
