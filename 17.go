/*
	Реализовать бинарный поиск встроенными методами языка.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binarySearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) >> 1
		midVal := arr[mid]

		if midVal < key {
			low = mid + 1
		} else if midVal > key {
			high = mid - 1
		} else {
			return mid // key found
		}
	}
	return -(low + 1) // key not found.
}

func generateSlice1(size int) []int {

	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) - rand.Intn(100)
	}
	return slice
}

func quicksort1(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort1(a[:left])
	quicksort1(a[left+1:])

	return a
}

func main() {
	slice := generateSlice1(100)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort1(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
	fmt.Println("\n--- Binary search key 44---\n")
	f := binarySearch(slice, 44)
	if f < 0 {
		fmt.Println("Not found")
	} else {
		fmt.Println(f)
	}
}
