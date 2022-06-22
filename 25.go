/*
	Реализовать собственную функцию sleep.
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func MySleep1(d time.Duration) {
	n := time.Now()
	n = n.Add(d)

	for time.Now().Before(n) {
		// Poor method so CPU is used
	}
}

func MySleep2(d time.Duration) {
	n := time.Now().Add(d)
	ctx, _ := context.WithDeadline(context.Background(), n)
	<-ctx.Done()
}

func main() {
	fmt.Println("Start sleep 1...")
	MySleep1(time.Second * 3)
	fmt.Println("Start sleep 2...")
	MySleep2(time.Second * 3)
	fmt.Println("Stop wait.")
}
