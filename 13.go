//13. Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 3
	b := 7
	fmt.Println(a, b)

	// by goland unknown as under cover its work
	a, b = b, a
	fmt.Println(a, b)

	// by ariphmetic
	b = b - a
	a = a + b
	b = a - b
	fmt.Println(a, b)

	// by bit operation
	a = a ^ b // a = 101^110 = 011
	b = a ^ b // b = 011^110 = 101
	a = a ^ b // a = 011^101 = 110
	fmt.Println(a, b)
}
