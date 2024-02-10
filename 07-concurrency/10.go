package main

import (
	"fmt"
	"time"
)

func main() {
	arr1 := []int{7, 2, 8, -9, 4, 0}
	arr2 := []int{17, 12, 18, -19, 14, 10}
	c := make(chan int)
	go sum(c, arr1)
	time.Sleep(time.Second * 2)
	go sum(c, arr2)
	x := <-c
	y := <-c

	fmt.Println(x, y)
}

func sum(c chan int, arr []int) {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	c <- sum
}
