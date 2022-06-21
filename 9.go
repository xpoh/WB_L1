/*
	9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
	массива, во второй — результат операции x*2, после чего данные из второго
	канала должны выводиться в stdout.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Create array
	var x [10]int
	for i, _ := range x {
		x[i] = rand.Int()
	}
	// create channels for in/out
	chIn := make(chan int)
	chOut := make(chan int)

	// create go for put data in array. After put channel close
	go func() {
		for i, _ := range x {
			chIn <- x[i]
		}
		close(chIn)
	}()

	// create go for get data from in channel and put x*2 values to output channel
	go func() {
		for x := range chIn {
			chOut <- x * 2
		}
		close(chOut)
	}()

	// read and print to console data while chan is open
	for x := range chOut {
		fmt.Println(x)
	}
}
