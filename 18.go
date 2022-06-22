/*
	18. Реализовать структуру-счетчик, которая будет инкрементироваться в
	конкурентной среде. По завершению программа должна выводить итоговое
	значение счетчика.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter interface {
	IncAtomic()
	IncMux()
	GetValue() int32
}

type Count struct {
	counter int32
	mux     sync.Mutex
}

func (c *Count) IncAtomic() {
	atomic.AddInt32(&c.counter, 1)
}

func (c *Count) IncMux() {
	c.mux.Lock()
	c.counter++
	c.mux.Unlock()
}

func (c *Count) GetValue() int32 {
	return c.counter
}

func main() {
	cnt := Count{}
	for i := 0; i < 100; i++ {
		go func() {
			cnt.IncAtomic()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("Atomic: ", cnt.GetValue())

	for i := 0; i < 100; i++ {
		go func() {
			cnt.IncMux()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("Mutex: ", cnt.GetValue())
}
