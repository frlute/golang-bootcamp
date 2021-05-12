package semaphore

import (
	"fmt"
	"testing"
	"time"
)

func Test_Semaphore(t *testing.T) {
	sem := NewSemaphore(3)

	doneC := make(chan bool, 1)
	totalProcess := 10

	for i := 1; i <= totalProcess; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			longRunningProcess(v)
			if v == totalProcess {
				doneC <- true
			}
		}(i)
	}
	for v := range doneC {
		if v {
			return
		}
	}
}

func longRunningProcess(taskID int) {
	fmt.Println(time.Now().Format("15:04:05"), "running tash with ID", taskID)
	time.Sleep(2 * time.Second)
}
