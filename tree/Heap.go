package tree

import "fmt"

// Heap _
type Heap struct {
	// 数组，从下标1开始存储数据
	items []int
	// 堆可以存储的最大数据个数
	capacity int
	// 堆中已经存储的数据个数
	length int
}

// NewHeap _
func NewHeap(capacity int) *Heap {
	return &Heap{
		items:    make([]int, capacity+1),
		capacity: capacity,
		length:   0,
	}
}

func (h *Heap) insert(item int) bool {
	if h.length == h.capacity {
		fmt.Println("the heap is full")
		return false
	}
	h.length++
	h.items[h.length] = item

	length := h.length
	for length/2 > 0 && h.items[length] > h.items[length/2] {
		// swap()函数作用：交换下标为i和i/2的两个元素
		swap(h.items, length, length/2)
		length = length / 2
	}

	return true
}

func swap(items []int, originPos, targetPos int) {
	items[originPos], items[targetPos] = items[targetPos], items[originPos]
}
