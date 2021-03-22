package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRetriable struct {
	count int
}

func (mr *MockRetriable) Exec() error {
	mr.count++
	if mr.count < 3 {
		return errors.New("test error")
	}
	return nil
}

func Test_New(t *testing.T) {
	type args struct {
		ret        Retrieable
		maxAttempt int
		loggerFunc func(error)
	}

	cases := []struct {
		name            string
		args            args
		expectedRetrier Retrier
		expectedErr     error
	}{
		{
			name: "normal no error",
			args: args{
				ret:        new(MockRetriable),
				maxAttempt: 5,
			},
			expectedRetrier: Retrier{
				attempt:     0,
				maxAttempt:  5,
				retrierable: new(MockRetriable),
			},
			expectedErr: nil,
		},
		{
			name: "error max attempt less than 1",
			args: args{
				ret: new(MockRetriable),
			},
			expectedRetrier: Retrier{
				retrierable: new(MockRetriable),
			},
			expectedErr: ErrMaxAttempt,
		},
		{
			name: "error retriable nil",
			args: args{
				maxAttempt: 1,
			},
			expectedRetrier: Retrier{},
			expectedErr:     ErrRetriableNil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := NewRetrier(c.args.ret, c.args.maxAttempt, c.args.loggerFunc)
			assert.Equal(tt, c.expectedErr, err)
			if r != nil {
				assert.Equal(tt, c.expectedRetrier, *r)
			}
		})
	}
}

func Test_do(t *testing.T) {
	r := Retrier{
		retrierable: new(MockRetriable),
	}

	r.do()
	assert.Equal(t, 1, r.attempt)
}

func Test_run(t *testing.T) {
	obj := new(MockRetriable)

	r := Retrier{
		maxAttempt:  3,
		retrierable: obj,
	}

	r.run()
	assert.Equal(t, 3, obj.count)
	assert.Equal(t, 3, r.attempt)
}
