/*
	Дана переменная int64. Разработать программу которая устанавливает i-й бит в
	1 или 0.

*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	v := rand.Int63() // random value
	nbit := uint64(1) // bit index
	vbit := uint64(0) // bit value

	fmt.Printf("v=%64b\n", v)
	fmt.Printf("set %v-th bit is %v\n", nbit, vbit)

	v1 := ^(uint64(1) << nbit) // create mask for low nbit
	v1 = v1 & uint64(v)        // low nbit
	v1 = v1 | (vbit << nbit)   // set nbit in vbit value
	v2 := int64(v1)
	fmt.Printf("v=%64b\n", v2)
}
