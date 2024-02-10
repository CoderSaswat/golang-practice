package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	maxRange := 100
	go checkGetEvenNos(maxRange, c, done)

	//with range keyword
	//for val := range c {
	//	fmt.Println(val)
	//}

	//with select keyword
	for {
		select {
		case val, ok := <-c:
			if !ok {
				fmt.Println("chanel is closed")
				return
			}
			fmt.Println(val)
		case <-done:
			fmt.Println("done")
		default:
			fmt.Println("default")
		}
	}
}

func checkGetEvenNos(maxRange int, c chan int, done chan bool) {
	for i := 1; i <= maxRange; i++ {
		if i%2 == 0 {
			c <- i
		}
		if i == i+1 {
			break
		}
	}
	done <- true
	close(c)
}
