package utils

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// 参考 env/go/pkg/mod/git.chengfayun.net/triones/compass@v1.11.2-0.20200709085637-15c956d9002a/utils/grace.go

// Grace _
type Grace struct {
	timeout  time.Duration
	cleanups []cleanFunc
	// 用来保护cleanups
	mutex sync.Mutex
	// 只调用Cleanup一次，调用后就不允许再执行Run。
	// 使用atomic.LoadUint32/StoreUint32保护
	done uint32
}

type cleanFunc struct {
	description string
	operation   func() error
}

// NewGrace _
func NewGrace(timeout time.Duration) *Grace {
	return &Grace{
		timeout: timeout,
		// cleanups: make([]cleanFunc),
	}
}

// Register 注册清理函数
// TODO 考虑清理函数顺序问题
func (g *Grace) Register(op string, f func() error) {
	if f == nil {
		panic("register nil cleanup function")
	}
	if atomic.LoadUint32(&g.done) == 1 {
		panic("register after the cleanup function")
	}

	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.cleanups = append(g.cleanups, cleanFunc{description: op, operation: f})
}

// Cleanup 公开这个方法，如果在最终启动起所有服务的过程中出错，可以调用Cleanup清理资源
// Cleanup只会执行一次注册过的清理函数，调用Cleanup后，再执行Cleanup/Register/Run都不会执行
func (g *Grace) Cleanup() {
	if !g.firstExecute() {
		return
	}

}

// Run _
func (g *Grace) Run() {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")
		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(g.timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit.", g.timeout.Microseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		g.execCleanups()

		close(wait)
	}()
}

func (g *Grace) execCleanups() {
	// 接下来会调用清理函数，阻止其他协程再执行Cleanup
	atomic.StoreUint32(&g.done, 1)

	wg := sync.WaitGroup{}
	// Do the operations asynchronously to save time
	// 反向调用，一般创建需要清理的资源时先创建的后关闭
	for index := len(g.cleanups) - 1; index >= 0; index-- {
		wg.Add(1)
		operation := g.cleanups[index]
		innerOp := operation.operation
		innerKey := operation.description
		go func() {
			defer wg.Done()
			log.Printf("cleaning up: %s", innerKey)
			if err := innerOp(); err != nil {
				log.Printf("%s: clean up failed: %s", innerKey, err.Error())
				return
			}

			log.Printf("%s was shutdown gracefully", innerKey)

		}()
	}

	wg.Wait()
}

func (g *Grace) firstExecute() bool {
	if atomic.LoadUint32(&g.done) == 1 {
		return false
	}
	return true
}
