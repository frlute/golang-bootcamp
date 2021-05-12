package semaphore

/*
Semaphore is a variable or abstract data type used to control access to a common resource by multiple processes in a concurrent system such as a multitasking operating system.(信号量是一种变量或抽象数据类型，用于控制并发系统（例如多任务操作系统）中的多个进程对公共资源的访问。)

使用场景：
- 限流
- 提高系统健壮性
*/

// Semaphore _
type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semC chan struct{}
}

// NewSemaphore create a semaphore instantiation
func NewSemaphore(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}
