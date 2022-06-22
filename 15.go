/*
	15. К каким негативным последствиям может привести данный фрагмент кода, и как
	это исправить? Приведите корректный пример реализации.

	var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}
	func main() {
		someFunc()
	}
	Проблема, что выделяется большой объем памяти, который не используется
	justString = v[:100] - здесь все равно выделяется новая память
*/
package main

var justString string

func createHugeString(int) string {
	return "hello"
}

func someFunc() {
	justString = createHugeString(100)
}
func main() {
	someFunc()
}
