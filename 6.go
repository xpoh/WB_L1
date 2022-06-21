/*
	Реализовать все возможные способы остановки выполнения горутины.
*/
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{} // waitgroup for wait for all work done

	// Stop by cancel func:
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)   // inc for one worker
	go func() { // run worker
		defer wg.Done()
		<-ctx.Done() // check for job done
		fmt.Println("Exit bu func cancel")
		return
	}()
	time.Sleep(time.Second)
	cancel()

	// Stop by Timeout:
	ctxt, _ := context.WithTimeout(context.Background(), time.Second) // context for workers with timeout
	wg.Add(1)                                                         // inc for one worker
	go func() {                                                       // run worker
		defer wg.Done()
		<-ctxt.Done() // check for job done
		fmt.Println("Exit by timeout")
		return
	}()

	// Stop by deadline:
	ctxd, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second)) // context for workers with deadline
	wg.Add(1)                                                                            // inc for one worker
	go func() {                                                                          // run worker
		defer wg.Done()
		<-ctxd.Done() // check for job done
		fmt.Println("Exit by deadline")
		return
	}()

	// Stop by signal notify:
	ch := make(chan os.Signal, 1)
	fmt.Println("Press ctrl+c")
	signal.Notify(ch, os.Interrupt)
	wg.Add(1)   // inc for one worker
	go func() { // run worker
		defer wg.Done()
		<-ch // check for job done
		fmt.Println("Exit by SIG_INT")
		return
	}()

	// Interrupt by stop main func:
	go func() { // run worker
		time.Sleep(time.Hour)
	}()

	// Interrupt self then job done
	go func() { // run worker
		// job done
	}()

	// Stop by self channel:
	exit := make(chan bool)
	wg.Add(1)   // inc for one worker
	go func() { // run worker
		defer wg.Done()
		<-exit // check for job done
		fmt.Println("Exit by self chan")
		return
	}()
	exit <- true
	wg.Wait()
}
