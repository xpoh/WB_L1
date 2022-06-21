package main

import (
	"fmt"
	"reflect"
)

// Способ 1
func PrintType1(any interface{}) {
	fmt.Println(reflect.TypeOf(any).String())
}

// Способ 2
func PrintType2(i interface{}) {
	// int, string, bool, channel
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan")
	case map[int]int:
		fmt.Println("map[int]int")
	}
}

func main() {
	a := true
	b := 10
	c := "Hello"
	d := make(chan int)
	e := make(map[int]int)

	PrintType1(a)
	PrintType1(b)
	PrintType1(c)
	PrintType1(d)
	PrintType1(e)

	PrintType2(a)
	PrintType2(b)
	PrintType2(c)
	PrintType2(d)
	PrintType2(e)

}
