/*
	24. Разработать программу нахождения расстояния между двумя точками, которые
	представлены в виде структуры Point с инкапсулированными параметрами x,y и
	конструктором.
*/
package main

import (
	"fmt"
	"math"
)

type Pointer interface {
	DistanceTo(p *Point) float64
}

type Point struct {
	x, y float64
}

func (pt Point) DistanceTo(p *Point) float64 {
	return math.Sqrt((pt.x-p.x)*(pt.x-p.x) + (pt.y-p.y)*(pt.y-p.y))
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func main() {
	p1 := NewPoint(-1, 3)
	p2 := NewPoint(2, -3.14)
	fmt.Println("Distance p1 to p2 =", p1.DistanceTo(p2))
}
