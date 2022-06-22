/*
	26. Разработать программу, которая проверяет, что все символы в строке
	уникальные (true — если уникальные, false etc). Функция проверки должна быть
	регистронезависимой.
	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/
package main

import (
	"fmt"
	"strings"
)

func StrUniq(str string) bool {
	m := make(map[rune]struct{})
	for _, s := range strings.ToUpper(str) {
		m[s] = struct{}{}
	}
	return len(m) == len(str)
}

func main() {
	fmt.Println("abcd", StrUniq("abcd"))
	fmt.Println("abCdefAaf", StrUniq("abCdefAaf"))
	fmt.Println("aabcd", StrUniq("aabcd"))
}
