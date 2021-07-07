package queue

// CirculationQueue _
type CirculationQueue struct {
	items    []interface{}
	capacity int
	head     int // 队头下标
	tail     int //队尾下标
}

// TODO 循环队列是不是用循环链表更好一些

// NewCirculationQueue _
func NewCirculationQueue(capacity int) *CirculationQueue {
	if capacity == 0 {
		return nil
	}
	return &CirculationQueue{
		items:    make([]interface{}, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

func (q *CirculationQueue) enqueue(item interface{}) bool {
	if (q.tail+1)%q.capacity == q.head {
		return false
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *CirculationQueue) dequeue() interface{} {
	// 如果head == tail 表示队列为空
	if q.head == q.tail {
		return nil
	}
	ret := q.items[q.head]
	q.head = (q.head + 1) % q.capacity
	return ret
}
