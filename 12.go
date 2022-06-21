/*
	12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
	собственное множество
*/

package main

import "fmt"

// В виду неопределенности постановки задачи, решимл выбрать уникальные элементы
func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{})

	for _, s := range str {
		m[s] = struct{}{}
	}
	for k, _ := range m {
		fmt.Println(k)
	}
}
