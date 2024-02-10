package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	//r := transformNumbers(doubleNumber, arr)
	r := transformNumbers(func(a int) int {
		return a * a * a
	}, arr)
	//operation := "triple"
	//r := transformNumbers(getOperationFn(operation), arr)

	//r := transformNumbers(createTransform(10), arr)

	fmt.Println(r)

}

func getOperationFn(operation string) func(a int) int {
	if operation == "double" {
		return doubleNumber
	}
	return tripleNumber
}

func transformNumbers(fn func(int) int, arr []int) []int {
	var r []int
	for _, num := range arr {
		r = append(r, fn(num))
	}
	return r
}

func doubleNumber(a int) int {
	return a * a
}
func tripleNumber(a int) int {
	return a * a * a
}

func createTransform(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}
