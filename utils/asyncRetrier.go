package utils

import "errors"

//AsyncRetrier 实现一种自动重试的统一封装(适用于请求发出并忘记的场景)
type AsyncRetrier struct {
	maxAttempt  int
	attempt     int
	done        bool
	retrierable Retrieable
	loggerFunc  func(error)
}

// Retrieable _
type Retrieable interface {
	Exec() error
}

// define some error message
var (
	ErrRetriableNil = errors.New("Retrieable object must not be nil")
	ErrMaxAttempt   = errors.New("Max attempt should be greater than zero")
)

// NewAsyncRetrier contstructing new AsyncRetrier object
// ret are retriable object that has Exec method
// maxAttempt is the maximum trying attempt before it ends
// loggerFunc is optional if you want to put logger of any error does occurs in the process
func NewAsyncRetrier(ret Retrieable, maxAttempt int, loggerFunc func(error)) (*AsyncRetrier, error) {
	if maxAttempt < 1 {
		return nil, ErrMaxAttempt
	}
	if ret == nil {
		return nil, ErrRetriableNil
	}

	return &AsyncRetrier{
		maxAttempt:  maxAttempt,
		attempt:     0,
		retrierable: ret,
		loggerFunc:  loggerFunc,
	}, nil
}

// Start startup retry flow
func (r *AsyncRetrier) Start() {
	go r.run()
}

func (r *AsyncRetrier) run() {
	for !r.isDone() {
		err := r.do()
		if err == nil {
			r.done = true
			return
		}
		if r.loggerFunc != nil {
			r.loggerFunc(err)
		}

	}
}

func (r *AsyncRetrier) do() error {
	r.attempt++
	return r.retrierable.Exec()
}

func (r *AsyncRetrier) isDone() bool {
	return r.attempt >= r.maxAttempt || r.done
}
