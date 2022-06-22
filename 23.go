/*
	23. Удалить i-ый элемент из слайса.
*/
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	i := 3

	// Способ 1
	b := a[:i]
	c := a[i+1:]
	for _, v := range c {
		b = append(b, v)
	}
	a = b
	fmt.Println(a)

	// Способ 2
	for j := i; j < len(a)-1; j++ {
		a[j] = a[j+1]
	}
	a = a[:len(a)-1]

	fmt.Println(a)

}
