package main

import "fmt"

func sum(n1 int, n2 int) int {
	return n1 + n2
}
func subtract(n1 int, n2 int) int {
	return n1 - n2
}
func operate(fn func(int, int) int, n1, n2 int) int {
	return fn(n1, n2)
}

func multiplyBy() func(int, int) int {
	return func(n1, n2 int) int {
		return n1 * n2
	}
}

func printConsole(name string) {
	fmt.Println("your name is :" + name)
}

func applyPrintConsole(fn func(string), name string) {
	fn(name)
}

func getClicker() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//result1 := operate(sum, 10, 5)
	//result2 := operate(subtract, 10, 5)
	//fmt.Println(result1)
	//fmt.Println(result2)

	//resultFn := multiplyBy()
	//fmt.Println(resultFn(10, 3))

	//result := getClicker()
	//fmt.Println(result())
	//fmt.Println(result())

	applyPrintConsole(printConsole, "saswat")
	applyPrintConsole(printConsole, "papun")
}
