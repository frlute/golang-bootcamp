package workerpool

import (
	"log"
)

/*
使用场景：在有限的 CPU 和 内存资源条件下，需要并发执行多个任务时，worker pool 则是一种在有限资源下最大化并发的方式
*/

// WorkerPool is a contract for Worker Pool implementation
type WorkerPool interface {
	Run()
	AddTask(task func())
	Completed() bool
}

type workerPool struct {
	maxWorker         int
	queuedTaskChannel chan func()
	// 标记是否已经开始处理 task
	started bool
}

// NewWorkerPool will create an instance of WorkerPool.
func NewWorkerPool(maxWorker int) WorkerPool {
	return &workerPool{
		maxWorker:         maxWorker,
		queuedTaskChannel: make(chan func()),
	}
}

func (wp *workerPool) Run() {
	for idx := 0; idx < wp.maxWorker; idx++ {
		wID := idx + 1
		log.Printf("[WorkerPool] Worker %d has been spawned", wID)
		go func(workerID int) {
			for task := range wp.queuedTaskChannel {
				log.Printf("[WorkerPool] Worker %d start processing task", workerID)
				task()
				log.Printf("[WorkerPool] Worker %d finish processing task", workerID)
				if !wp.started {
					wp.started = true
				}
			}
		}(wID)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskChannel <- task
}

func (wp *workerPool) totalQueuedTask() int {
	return len(wp.queuedTaskChannel)
}

func (wp *workerPool) Completed() bool {
	return wp.started && wp.totalQueuedTask() == 0
}
