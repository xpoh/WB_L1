/*
	19. Разработать программу, которая переворачивает подаваемую на ход строку
	(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
package main

import "fmt"

func ReverseString(str string) string {
	var r string
	for _, c := range str {
		r = string(c) + r
	}
	return r
}

func main() {
	str := "12345"
	fmt.Println("String =", str)
	fmt.Println("Reverse string =", ReverseString(str))
}
