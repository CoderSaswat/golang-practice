package main

import "fmt"

func main() {
	//result := add(1, 2)
	//result := add(1.6, 2.4)
	result := add("saswat ", "panda")
	fmt.Println(result)
}
func add[T int | float64 | string](a, b T) T {
	return a + b
}
