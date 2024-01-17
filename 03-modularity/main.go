package main

import (
	"fmt"
	"my-app/calculator"
)

func main() {
	//fmt.Println("hi modularity")
	fmt.Println(calculator.Add(10, 20))
	fmt.Println(calculator.Subtract(20, 10))
	//fmt.Println(utils.IsEven(2))
	fmt.Println("Operation Count :", calculator.OpCount())

}
