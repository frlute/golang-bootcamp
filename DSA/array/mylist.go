package array

import "fmt"

type MyArray struct {
	items    []int
	length   int
	capacity int
}

func newMyArray(length int, capacity ...int) Array {
	var items []int
	cap := 0
	if len(capacity) == 1 {
		cap = capacity[0]
		items = make([]int, 0, cap)
	} else {
		items = make([]int, 0)
	}

	return &MyArray{
		length:   0,
		items:    items,
		capacity: cap,
	}
}

func (l *MyArray) Append(value ...int) []int {
	m := l.length
	n := m + len(value)
	// go 切片的动态扩容
	// 使用 copy 因为它复制可以处理共享相同基础数组的源和目标片，正确地处理重叠片.
	// TODO 此处内部使用 slice 作为数组, 故 增加 = 条件，防止数组越界
	if n >= l.capacity {
		// 避免 n=0, allocate double what's needed, for future growth.
		l.capacity = (l.length + 1) * 2
		newSlice := make([]int, n, l.capacity)
		copy(newSlice, l.items)

		l.items = newSlice
	}

	l.items = l.items[0:n]
	copy(l.items[m:n], value)
	l.length += len(value)
	return l.items
}

// Insert 插入指定位置的 value, 暂不考虑 cap 问题
func (l *MyArray) Insert(index, value int) {
	if index > l.length-1 {
		fmt.Printf("insert pos %d out of the array length %d\n", index, l.length)
		return
	}
	tmp := make([]int, 0, l.length+1)
	tmp = append(tmp, l.items[0:index]...)
	tmp = append(tmp, value)
	tmp = append(tmp, l.items[index:]...)
	l.items = tmp
	l.length++
}

// GetIndex 返回第一个匹配到索引
func (l *MyArray) GetIndex(target int) int {
	for index, value := range l.items {
		if value == target {
			return index
		}
	}

	return -1
}

func (l *MyArray) Remove(target int) {
	index := l.GetIndex(target)
	if index != -1 {
		tmp := make([]int, 0, l.length-1)
		tmp = append(tmp, l.items[0:index]...)
		tmp = append(tmp, l.items[index+1:]...)
		l.items = tmp
		l.length--
	}
}

func (l *MyArray) Cap() int {
	return l.capacity
}

func (l *MyArray) Len() int {
	return l.length
}

func (l *MyArray) Display() {
	fmt.Printf("%#v\n", l.items)
}
