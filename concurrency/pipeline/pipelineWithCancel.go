package pipeline

//ExecutorWithCancel _
// type ExecutorWithCancel func(<-chan struct{}, interface{}) (<-chan interface{}, error)

// // PipelineWithCancel _
// type PipelineWithCancel interface {
// 	Pipe(executors ...ExecutorWithCancel) Pipeline
// 	Merge(<-chan struct{}) <-chan interface{}
// 	// Cancel(<-chan struct{}) // 支持取消
// }

// type pipelineWithCancel struct {
// 	// dataC     chan interface{}
// 	// errC      chan error
// 	executors []ExecutorWithCancel
// 	// done      chan struct{} // 取消信号
// }

// // NewPipelineWithCancel _
// func NewPipelineWithCancel(f func(...interface{}) <-chan interface{}) PipelineWithCancel {
// 	// 初始化处理
// 	f()

// 	return &pipelineWithCancel{
// 		executors: []ExecutorWithCancel{},
// 		// done:      make(chan struct{}),
// 	}
// }

// func (p *pipelineWithCancel) Pipe(executors ...Executor) Pipeline {
// 	// p.executors = append(p.executors, executors...)

// 	return p
// }

// func (p *pipelineWithCancel) Merge(done <-chan struct{}) <-chan interface{} {
// 	// var wg sync.WaitGroup
// 	out := make(chan interface{})

// 	// output := func(c <-chan interface{}) {
// 	// 	for n := range c {
// 	// 		out <- n
// 	// 	}
// 	// 	wg.Done()
// 	// }

// 	// wg.Add(len(p.executors))
// 	// for _, execute := range p.executors {
// 	// 	go output(execute)
// 	// }

// 	return out
// }

// func (p *pipepipelineWithCancel)run(inC <-chan interface{}, execute ExeExecutor) <-chan interface{} {
// 	for input := range inC {

// 	}

// }

// func (p *pipelineWithCancel) Cancel(done <-chan struct{}) {
// 	p.done <- <-done
// }
