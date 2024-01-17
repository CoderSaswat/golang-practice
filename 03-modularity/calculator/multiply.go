package calculator

import "fmt"

func init() {
	fmt.Println("calculator initialized - [multiply.go]")
}

func Multiply(x, y int) int {
	// opCount++
	opCount["multiply"]++
	println(opCount)
	return x * y
}
