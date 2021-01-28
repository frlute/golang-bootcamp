package queue

import "fmt"

// ArrayQueue _
type ArrayQueue struct {
	items    []string
	capacity int
	head     int // 队头下标
	tail     int //队尾下标
}

// NewArrayQueue _
func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		items:    make([]string, capacity),
		capacity: capacity,
	}
}

func (q *ArrayQueue) enqueue(item string) bool {
	// 队满
	if q.tail == q.capacity {
		// tail ==n && head==0，表示整个队列都占满了
		if q.head == 0 {
			fmt.Printf("已经最大容量, capacity %d,  items: %+v\n", q.capacity, q.items)
			return false
		}
		// 数据搬移
		for index := q.head; index < q.tail; index++ {
			q.items[index-q.head] = q.items[index]
		}
		// 搬移完之后重新更新head和tail
		q.tail -= q.head
		q.head = 0
	}
	q.items[q.tail] = item
	q.tail++

	return true
}

func (q *ArrayQueue) dequeue() string {
	// 表示空队列
	if q.head == q.tail {
		return ""
	}

	ret := q.items[q.head]
	q.items[q.head] = ""
	q.head++
	return ret
}
