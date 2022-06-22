/*
	22. Разработать программу, которая перемножает, делит, складывает, вычитает две
	числовых переменных a,b, значение которых > 2^20
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.Int{}
	b := big.Int{}
	c := big.Int{}

	var stra string = "782346547365876324785634783678567826347857823658726347865783465786347865287365"
	var strb string = "5847857487583475834785734875834775384573485"

	//fmt.Println("Input a:")
	//fmt.Scanln(&stra)
	//fmt.Println("Input b:")
	//fmt.Scanln(&strb)

	a.SetString(stra, 10)
	b.SetString(strb, 10)

	c = *a.Add(&a, &b)
	fmt.Println("Result a + b:", c.String())

	c = *a.Sub(&a, &b)
	fmt.Println("Result a - b:", c.String())

	c = *a.Mul(&a, &b)
	fmt.Println("Result a*b:", c.String())

	c = *a.Div(&a, &b)
	fmt.Println("Result a/b:", c.String())

}
