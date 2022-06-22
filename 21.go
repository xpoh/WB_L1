/*
	21. Реализовать паттерн «адаптер» на любом примере.
*/
package main

import (
	"fmt"
	"strings"
)

func ReverseWords(strs []string) []string {
	var result []string
	for i := range strs {
		result = append(result, strs[len(strs)-i-1])
	}
	return result
}

type StringToSliceAdapter interface {
	ReverseString(str string) string
}

func (i *InMemoryAdapter) ReverseString(str string) string {
	var r string
	sl := strings.Split(str, " ")
	sl = ReverseWords(sl)
	for _, s := range sl {
		r = r + " " + s
	}
	return r[1:]
}

type InMemoryAdapter struct{}

func main() {
	ia := InMemoryAdapter{}
	str := "12345 111111 2222222"
	fmt.Println("String =", str)
	fmt.Println("Reverse word =", ia.ReverseString(str))
}
