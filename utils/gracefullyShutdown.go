package utils

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
几个关键点：

1. Listen for the termination signal/s from the process manager like SIGTERM
2. We block the main function until the signal is received
3. After we received the signal, we will do clean-ups on our app and wait until those clean-up operations are done.
4. We also need to have a timeout to ensure that the operation won’t hang up the system.
*/

// operation is a clean up function on shutting down
type operation func(ctx context.Context) error

// gracefulShutdown waits for termination syscalls and doing clean up operations after received it
func gracefuleShutdow(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")
		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit.", timeout.Microseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		wg := sync.WaitGroup{}

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()
				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)

			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}

/*
使用case
func main() {
  // initialize some resources
  // e.g:
  // db, err := database.Initialize()
  // server, err := http.Initialize()

  // wait for termination signal and register database & http server clean-up operations
  wait := gracefulShutdown(context.Background(), 2 * time.Second, map[string]operation{
    "database": func(ctx context.Context) error {
      return db.Shutdown()
		},
    "http-server": func(ctx context.Context) error {
      return server.Shutdown()
		},
	// Add other cleanup operations here
  })

  <-wait
*/
