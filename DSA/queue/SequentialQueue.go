package queue

import (
	"fmt"
	"sync"
)

// 顺序队列
type sequentialQueue struct {
	items    []interface{}
	capacity int
	// the queue front pointer, 队头下标
	head int
	// the queue rear pointer, 队尾下标
	tail int
	sync.Mutex
}

// NewSequentialQueue _
func NewSequentialQueue(capacity int) Queue {
	return newSequentialQueueArray(capacity)
}

func newSequentialQueueArray(capacity int) *sequentialQueue {
	return &sequentialQueue{
		items:    make([]interface{}, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

func (q *sequentialQueue) Enqueue(item interface{}) {
	// 队满
	if q.IsFull() {
		// tail == n && head==0，表示整个队列都占满了
		if q.head == 0 {
			fmt.Printf("已经最大容量, capacity %d,  items: %#v\n", q.capacity, q.items)
			return
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
}

func (q *sequentialQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	ret := q.items[q.head]
	// 每次迁移效率很低，为了提高效率在入队达到最大容量时再迁移
	// q.items[q.head] = nil
	q.head++
	return ret
}

func (q *sequentialQueue) Peek() interface{} {
	return q.items[q.tail]
}

func (q *sequentialQueue) IsFull() bool {
	return (q.tail - q.head) == q.capacity
}

func (q *sequentialQueue) IsEmpty() bool {
	return q.head < 0 || q.head >= q.tail
}

func (q *sequentialQueue) Max() interface{} {
	max := q.items[q.head]

	for i := q.head + 1; i <= q.tail; i++ {
		if q.items[i] != nil && q.items[i].(int) > max.(int) {
			max = q.items[i]
		}
	}

	return max
}
