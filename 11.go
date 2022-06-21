/*
	11. Реализовать пересечение двух неупорядоченных множеств.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	set1 := make([]uint8, 100)
	set2 := make([]uint8, 100)

	var muxSet []uint8

	// Fill test date in set1,2
	for i, _ := range set1 {
		set1[i] = uint8(rand.Uint32())
	}
	for i, _ := range set2 {
		set2[i] = uint8(rand.Uint32())
	}
	// Find mux set (difficulty O(n*n)
	for _, v1 := range set1 {
		for _, v2 := range set2 {
			if v1 == v2 {
				muxSet = append(muxSet, v1)
				break
			}
		}
	}

	// Output all set
	fmt.Println("Set1=", set1)
	fmt.Println("Set2=", set2)
	fmt.Println("MuxSet=", muxSet)
}
