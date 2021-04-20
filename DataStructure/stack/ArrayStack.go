package stack

import "fmt"

// ArrayStack _
type ArrayStack struct {
	items           []string
	count           int  // 栈中元素个数
	size            int  // 栈大小
	dynamicCapacity bool // 是否动态扩容
}

// NewArrayStack _
func NewArrayStack(size int, dynamicCapacity bool) *ArrayStack {
	as := &ArrayStack{
		// TODO 应该定义为数组
		items: make([]string, size, size),
		count: 0,
		size:  size,
	}
	return as
}

func (s *ArrayStack) push(item string) bool {
	// 数组空间不够了，直接返回false，入栈失败。
	if s.count == s.size {
		if !s.dynamicCapacity {
			fmt.Printf("已经最大容量, count %d, size %d, items: %+v\n", s.count, s.size, s.items)
			return false
		}
		// 进行动态扩容
	}
	// TODO 注意并发
	// s.items = append(s.items, item)
	s.items[s.count] = item
	s.count++

	return true
}

func (s *ArrayStack) pop() string {
	// 栈为空，则直接返回null
	if s.count == 0 {
		return ""
	}
	// 返回下标为count-1的数组元素，并且栈中元素个数count减一
	item := s.items[s.count-1]
	s.count--
	s.items[s.count-1] = ""
	fmt.Printf("test, count %d, size %d, items: %+v", s.count, s.size, s.items)

	return item
}
