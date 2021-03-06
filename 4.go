/*
	Реализовать постоянную запись данных в канал (главный поток). Реализовать
	набор из N воркеров, которые читают произвольные данные из канала и
	выводят в stdout. Необходима возможность выбора количества воркеров при
	старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
	способ завершения работы всех воркеров.
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
)

func main() {
	N := 10

	ct := context.Background() // context for workers cancel
	ctx, cancel := context.WithCancel(ct)
	wg := sync.WaitGroup{} // waitgroup for wait for all work done

	ch := make(chan int)                               // main channel for random data
	exit := make(chan os.Signal, 1)                    // chan for SIGINT (ctrl+c)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM) // notify for ctrl+c

	for i := 0; i < N; i++ {
		wg.Add(1)   // inc for one worker
		go func() { // run worker
			defer wg.Done()
			for {
				select {
				case <-ctx.Done(): // check for job done
					fmt.Println("Terminate worker.")
					return
				case c := <-ch: // check data in chan
					fmt.Println(c)
				}
			}
		}()
	}

	for {
		select {
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
