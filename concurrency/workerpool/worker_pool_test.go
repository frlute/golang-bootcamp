package workerpool

import (
	"log"
	"runtime"
	"testing"
	"time"
)

func Test_WorkerPool(t *testing.T) {
	// For monitoring purpose.
	waitC := make(chan bool)
	go func() {
		for {
			log.Printf("[main] Total current goroutine: %d", runtime.NumGoroutine())
			time.Sleep(500 * time.Microsecond)
		}
	}()

	// Start Worker Pool.
	totalWorker := 5
	wp := NewWorkerPool(totalWorker)
	wp.Run()

	type result struct {
		id    int
		value int
	}

	totalTask := 10
	resultC := make(chan result, totalTask)

	for i := 0; i < totalTask; i++ {
		id := i + 1
		wp.AddTask(func() {
			log.Printf("[main] Starting task %d", id)
			time.Sleep(1 * time.Second)
			resultC <- result{id, id * 2}
		})
	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("[main] Task [%d] has been finished with result [%d]", res.id, res.value)
	}

	for !wp.Completed() {
		<-waitC
	}
}
