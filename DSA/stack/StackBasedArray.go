package stack

import (
	"fmt"
	"sync"
)

// StackBasedArray _
type StackBasedArray struct {
	items    []interface{}
	top      int64 // 数组中元素顶点位置，从 0 开始
	capacity int64 // 栈容量
	sync.Mutex
	// dynamic bool  // 是否动态扩容
}

// Options 栈创建可选参数
type Options struct {
	dynamic bool
}

// NewStackBasedArray _
func NewStackBasedArray(capacity int64) Stack {
	return &StackBasedArray{
		items:    make([]interface{}, capacity),
		top:      0,
		capacity: capacity,
	}
}

// Push Pushing (storing) an element on the stack.
func (s *StackBasedArray) Push(item interface{}) {
	// 数组空间不够了，直接返回false，入栈失败。
	if s.IsFull() {
		fmt.Println("Stack is full.")
		return
		// ("已经最大容量, count %d, size %d, items: %+v\n", s.length, s.capacity, s.items)
		// return false
		// if !s.dynamicCapacity {
		// 	fmt.Printf("已经最大容量, count %d, size %d, items: %+v\n", s.length, s.size, s.items)
		// 	return false
		// }
		// 进行动态扩容
	}
	// TODO 注意并发
	// s.items = append(s.items, item)
	s.items[s.top] = item
	s.top++

	// return true
}

// Pop Removing an element from the stack.
func (s *StackBasedArray) Pop() interface{} {
	// 栈为空，则直接返回null
	if s.IsEmpty() {
		fmt.Println("Could not retrieve data, Stack is empty.")
		return nil
	}

	// 返回下标为count-1的数组元素，并且栈中元素个数count减一
	// item := s.items[s.top-1]
	// s.items[s.top-1] = nil
	// s.top--
	// 改进方式，顺序栈可以不进行实际的 element 删除，只是移动 top 的位置向下移动一位
	item := s.items[s.top-1]
	s.top--

	return item
}

// Peek get the top data element of the stack, without removing it.
func (s *StackBasedArray) Peek() interface{} {
	return s.items[s.top-1]
}

// IsFull check if stack is full.
func (s *StackBasedArray) IsFull() bool {
	return s.top == s.capacity
}

// IsEmpty check if stack is empty.
func (s *StackBasedArray) IsEmpty() bool {
	return s.top == 0
}

// Flush _
func (s *StackBasedArray) Flush() {
	s.top = 0
}

// Display _
func (s *StackBasedArray) Display() {
	fmt.Printf("Stack items: %+v\n", s.items)
}
