package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 10
	c <- 12
	fmt.Println(<-c)
	fmt.Println(<-c)
}
