package stack

import "fmt"

// ArrayStack _
type ArrayStack struct {
	items    []interface{}
	count    int64 // 栈中元素个数
	capacity int64 // 栈大小
	// dynamicCapacity bool  // 是否动态扩容
}

// NewArrayStack _
func NewArrayStack(capacity int64, dynamicCapacity bool) Stack {
	return &ArrayStack{
		// TODO 应该定义为数组
		items:    make([]interface{}, capacity, capacity),
		count:    0,
		capacity: capacity,
	}
}

func (s *ArrayStack) Push(item interface{}) bool {
	// 数组空间不够了，直接返回false，入栈失败。
	if s.count == s.capacity {
		fmt.Printf("已经最大容量, count %d, size %d, items: %+v\n", s.count, s.capacity, s.items)
		return false
		// if !s.dynamicCapacity {
		// 	fmt.Printf("已经最大容量, count %d, size %d, items: %+v\n", s.count, s.size, s.items)
		// 	return false
		// }
		// 进行动态扩容
	}
	// TODO 注意并发
	// s.items = append(s.items, item)
	s.items[s.count] = item
	s.count++

	return true
}

func (s *ArrayStack) Pop() interface{} {
	// 栈为空，则直接返回null
	if s.count == 0 {
		return ""
	}
	// 返回下标为count-1的数组元素，并且栈中元素个数count减一
	item := s.items[s.count-1]
	s.items[s.count-1] = ""
	s.count--
	fmt.Printf("test, count %d, size %d, items: %+v", s.count, s.capacity, s.items)

	return item
}

func (s *ArrayStack) Len() int64 {
	return s.count
}

func (s *ArrayStack) Cap() int64 {
	return s.capacity
}
