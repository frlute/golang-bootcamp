package pipeline

import (
	"fmt"
	"testing"
	"time"
)

func Test_pipeline_Pipe(t *testing.T) {
	type fields struct {
		executors []Executor
	}

	tests := []struct {
		name   string
		fields fields
		check  func(f fields) bool
	}{
		{
			name: "should success adding executor to executors",
			fields: fields{
				executors: []Executor{
					func(in interface{}) (interface{}, error) {
						return 1, nil
					},
				},
			},
			check: func(f fields) bool {
				return len(f.executors) == 1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pipeline{
				dataC:     nil,
				errC:      nil,
				executors: nil,
			}

			p.Pipe(tt.fields.executors...)
			if !tt.check(fields{p.executors}) {
				t.Errorf("pipeline.Pipe() not run as expected")
			}
		})
	}
}

func Test_pipeline(t *testing.T) {
	N := 5
	startTime := time.Now()

	for i := 0; i < N; i++ {
		result := addFoo(addQuoute(square(multiplyTwo(i))))
		fmt.Printf("Result: %s\n", result)
	}
	fmt.Printf("Elapsed time without concurrency: %s\n", time.Since(startTime)) // ~20 seconds

	p := NewPipeline(func(inC chan interface{}) {
		defer close(inC)
		for i := 0; i < N; i++ {
			inC <- i
		}
	})
	outC := p.Pipe(
		func(in interface{}) (interface{}, error) {
			return multiplyTwo(in.(int)), nil
		},
		func(in interface{}) (interface{}, error) {
			return square(in.(int)), nil
		},
		func(in interface{}) (interface{}, error) {
			return addQuoute(in.(int)), nil
		},
		func(in interface{}) (interface{}, error) {
			return addFoo(in.(string)), nil
		},
	).Merge()

	startTimeC := time.Now()
	for result := range outC {
		fmt.Printf("pipline Result: %s\n", result)
	}
	fmt.Printf("Elapsed time with concurrency by pipeline: %s", time.Since(startTimeC)) // ~16 seconds
}

func multiplyTwo(v int) int {
	time.Sleep(1 * time.Second)
	return v * 2
}

func square(v int) int {
	time.Sleep(1 * time.Second)
	return v * v
}

func addQuoute(v int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("'%d'", v)
}

func addFoo(v string) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("%s - Foo", v)
}
