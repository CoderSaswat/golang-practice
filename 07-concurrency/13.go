package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second * 3)
	boom := time.After(time.Second * 10)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Println("default")
		}
	}
}
