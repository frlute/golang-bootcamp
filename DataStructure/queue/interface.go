package queue

type Queue interface {
	Enqueue(item interface{}) bool
	Dequeue() interface{}
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
