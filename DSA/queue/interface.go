package queue

type Queue interface {
	//  add (store) an item to the queue.
	Enqueue(item interface{})
	// remove (access) an item from the queue.
	Dequeue() interface{}
	// Gets the element at the front of the queue without removing it.
	Peek() interface{}
	IsFull() bool
	IsEmpty() bool

	Display()
}

// Deque 双端队列
type Deque interface {
	// 头部插入与删除
	Push(item interface{}) bool
	Pop() interface{}

	// 尾部插入与删除
	Inject(item interface{}) bool
	Eject() interface{}
}
