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
	N := 5 * time.Second

	ct := context.Background() // context for workers cancel
	ctx, _ := context.WithTimeout(ct, N)
	wg := sync.WaitGroup{} // waitgroup for wait for all work done

	ch := make(chan int)                 // main channel for random data
	exit := make(chan os.Signal)         // chan for sig_term (ctrl+c)
	signal.Notify(exit, syscall.SIGTERM) // notify for ctrl+c

	wg.Add(1)   // inc for one worker
	go func() { // run worker
		for {
			select {
			case <-ctx.Done(): // check for job done
				fm
				break
			case c := <-ch: // check data in chan
				fmt.Println(c)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done(): // check SIG_TERM
			ctx.Done() // terminate all workers
			wg.Wait()  // wait while terminate
			break
		default:
			ch <- rand.Int() // put random data in chan
		}
	}
}
