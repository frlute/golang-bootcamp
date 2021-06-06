package pipeline

import (
	"fmt"
	"testing"
	"time"
)

func Test_pipelineWithCancel(t *testing.T) {
	// N := 5

	// initFunc := func(items ...interface{}) <-chan interface{} {
	// 	out := make(chan interface{}, len(items))
	// 	for _, n := range items {
	// 		out <- n
	// 	}
	// 	close(out)
	// 	return out
	// }

	// p := NewPipelineWithCancel(initFunc)
	// outC := p.Pipe(
	// 	func(in interface{}) (interface{}, error) {
	// 		return multiplyTwo(in.(int)), nil
	// 	},
	// 	func(in interface{}) (interface{}, error) {
	// 		return square(in.(int)), nil
	// 	},
	// 	func(in interface{}) (interface{}, error) {
	// 		return addQuoute(in.(int)), nil
	// 	},
	// 	func(in interface{}) (interface{}, error) {
	// 		return addFoo(in.(string)), nil
	// 	},
	// ).Merge()

	// startTimeC := time.Now()
	// for result := range outC {
	// 	fmt.Printf("pipline Result: %s\n", result)
	// }
	// fmt.Printf("Elapsed time with concurrency by pipeline: %s", time.Since(startTimeC)) // ~16 seconds
}

func multiplyTwoWithCancel(v int) int {
	time.Sleep(1 * time.Second)
	return v * 2
}

func squareWithCancel(v int) int {
	time.Sleep(1 * time.Second)
	return v * v
}

func addQuouteWithCancel(v int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("'%d'", v)
}

func addFooWithCancel(v string) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("%s - Foo", v)
}
