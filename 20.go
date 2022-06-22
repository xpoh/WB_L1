/*
	20. Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"fmt"
	"strings"
)

func ReverseWord(str string) string {
	var r string
	sl := strings.Split(str, " ")

	for _, s := range sl {
		r = s + " " + r
	}
	return r[0 : len(r)-1]
}

func main() {
	str := "12345 111111 2222222"
	fmt.Println("String =", str)
	fmt.Println("Reverse word =", ReverseWord(str))
}
