package queue

import "fmt"

// QueueBasedArray _
type queueBasedArray struct {
	items    []interface{}
	capacity int64
	head     int64 // 队头下标
	tail     int64 //队尾下标
}

// NewQueueBasedArray _
func NewQueueBasedArray(capacity int64) Queue {
	return newQueueBasedArrayArray(capacity)
}

func newQueueBasedArrayArray(capacity int64) *queueBasedArray {
	return &queueBasedArray{
		// TODO 前期应该分配较小的内存，之后按需动态扩展，但是不能超出 capacity
		items:    make([]interface{}, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

func (q *queueBasedArray) Enqueue(item interface{}) bool {
	// 队满
	if q.tail == q.capacity {
		// tail ==n && head==0，表示整个队列都占满了
		if q.head == 0 {
			fmt.Printf("已经最大容量, capacity %d,  items: %#v\n", q.capacity, q.items)
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
	// TODO 加锁防止并发引起的竞争
	q.items[q.tail] = item
	q.tail++

	return true
}

func (q *queueBasedArray) Dequeue() interface{} {
	// 表示空队列
	if q.head == q.tail {
		return ""
	}

	ret := q.items[q.head]
	// 每次迁移效率很低，为了提高效率在入队时达到最大容量时再迁移
	q.items[q.head] = nil
	q.head++
	return ret
}

func (q *queueBasedArray) Max() interface{} {
	max := q.items[q.head]

	for i := q.head + 1; i <= q.tail; i++ {
		if q.items[i] != nil && q.items[i].(int) > max.(int) {
			max = q.items[i]
		}
	}

	return max
}
