package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius int
}

type Square struct {
	side int
}

func (circle Circle) area() float64 {
	return float64(circle.radius) * float64(circle.radius) * math.Pi
}

func (circle Circle) String() string {
	return "my formatting implementation"
}

func (square Square) perimeter() float64 {
	return float64(square.side * 4)
}

func (square Square) area() float64 {
	return float64(square.side) * float64(square.side)
}
func (circle Circle) perimeter() float64 {
	return 2 * float64(circle.radius) * math.Pi
}

func showInformation(shape Shape) {
	fmt.Printf("area %v\n", shape.area())
	fmt.Printf("perimeter %v\n", shape.perimeter())
}

func main() {
	s1 := Square{side: 10}
	//var shape1 Shape
	//shape1 = s1
	c1 := Circle{radius: 10}
	//var shape2 Shape
	//shape2 = c1
	//showInformation(s1)
	//showInformation(c1)

	fmt.Println(c1)
	fmt.Println(s1)
}
