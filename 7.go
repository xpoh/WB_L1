/*
	Реализовать конкурентную запись данных в map
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	mux := sync.Mutex{} // use mutext for concurency write to map
	m := make(map[int]string)

	for i := 0; i < 10; i++ {
		go func() {
			i := rand.Int()
			mux.Lock()
			m[i] = strconv.Itoa(i)
			mux.Unlock()
		}()
	}
	time.Sleep(time.Second)
	for i, s := range m {
		fmt.Printf("m[%v]=%v\n", i, s)
	}

}
