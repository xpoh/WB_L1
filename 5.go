/*
	5. Разработать программу, которая будет последовательно отправлять значения в
	канал, а с другой стороны канала — читать. По истечению N секунд программа
	должна завершаться
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	N := 10 * time.Second

	ct := context.Background()                // context for workers cancel
	ctx, cancel := context.WithTimeout(ct, N) // Set timout
	wg := sync.WaitGroup{}                    // waitgroup for wait for all work done

	ch := make(chan int)                               // main channel for random data
	exit := make(chan os.Signal, 1)                    // chan for SIGINT (ctrl+c)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM) // notify for ctrl+c

	wg.Add(1)   // inc for one worker
	go func() { // run worker
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): // check for job done
				fmt.Println("Terminate worker by timeout")
				return
			case c := <-ch: // check data in chan
				fmt.Println(c)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done(): // check SIGINT
			fmt.Println("Stop main by timeout")
			wg.Wait() // wait while terminate
			return
		case sig := <-exit: // check SIGINT
			fmt.Println(sig)
			cancel()  // terminate all workers
			wg.Wait() // wait while terminate
			return
		default:
			ch <- rand.Int() // put random data in chan
		}
	}
}
