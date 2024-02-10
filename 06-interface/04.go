package main

import "fmt"

type Shape2 interface {
	getArea() float64
}

type Triangle2 struct {
	base   float64
	height float64
}
type Square2 struct {
	sideLength float64
}

func (s Square2) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t Triangle2) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(shape Shape2) {
	fmt.Println("area is", shape.getArea())
}

func main() {
	sqaure := Square2{sideLength: 10}
	triangle := Triangle2{
		base:   10.0,
		height: 5.0,
	}
	printArea(triangle)
	printArea(sqaure)
}
