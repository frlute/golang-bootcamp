package pipeline

/*
Collection pipelines are a programming pattern where you organize some computation as a sequence of operations which compose by taking a collection as output of one operation and feeding it into the next.
		-- Collection Pipeline by Martin Fowler:
*/

//Executor _
type Executor func(interface{}) (interface{}, error)

// Pipeline _
type Pipeline interface {
	Pipe(executors ...Executor) Pipeline
	Merge() <-chan interface{}
}

type pipeline struct {
	dataC     chan interface{}
	errC      chan error
	executors []Executor
}

// NewPipeline _
func NewPipeline(f func(chan interface{})) Pipeline {
	inC := make(chan interface{})

	// 初始化处理
	go f(inC)

	return &pipeline{
		dataC:     inC,
		errC:      make(chan error),
		executors: []Executor{},
	}
}

func (p *pipeline) Pipe(executors ...Executor) Pipeline {
	p.executors = append(p.executors, executors...)

	return p
}

func (p *pipeline) Merge() <-chan interface{} {
	for _, execute := range p.executors {
		p.dataC, p.errC = p.run(p.dataC, execute)
	}

	return p.dataC
}

func (p *pipeline) run(inC <-chan interface{}, f Executor) (chan interface{}, chan error) {
	outC := make(chan interface{})
	errC := make(chan error)

	go func() {
		defer close(outC)
		for v := range inC {
			res, err := f(v)
			if err != nil {
				errC <- err
				continue
			}

			outC <- res
		}
	}()

	return outC, errC
}
