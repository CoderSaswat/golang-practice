package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//start := time.Now()
	//names := []string{"virat", "rohit", "rahul", "surya", "bumra"}
	//for _, name := range names {
	//	go printNames(name)
	//}
	//time.Sleep(time.Second * 2)

	//defer func() {
	//	fmt.Println(time.Since(start))
	//}()
	//name := "virat"
	//nameChnl := make(chan bool, 2)
	//go printNames(name, nameChnl)
	//fmt.Println(<-nameChnl)

	//nameChnl := make(chan string, 1)
	//nameChnl <- "rohit"
	//fmt.Println(<-nameChnl)
	//nameChnl <- "virat"
	//fmt.Println(<-nameChnl)
	//nameChnl <- "dhoni"
	//fmt.Println(<-nameChnl)
	//nameChnl <- "kl"
	//fmt.Println(<-nameChnl)

	//nameChnl := make(chan string)
	//go func() {
	//	fmt.Println(<-nameChnl)
	//}()
	//nameChnl <- "rohit"
	//close(nameChnl)

	//go func() {
	//	nameChnl <- "rohit"
	//}()
	//fmt.Println(<-nameChnl)
	//time.Sleep(time.Second * 2)

	nameChnl := make(chan string)
	//go func() {
	//	for value := range nameChnl {
	//		time.Sleep(time.Second * 1)
	//		fmt.Println(value)
	//	}
	//}()
	//
	//nameChnl <- "virat"
	//nameChnl <- "rohit"
	//nameChnl <- "pandya"
	//close(nameChnl)

	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go processNames(nameChnl, &wg)
	//// Send names into the unbuffered channel
	//nameChnl <- "virat"
	//nameChnl <- "rohit"
	//nameChnl <- "pandya"
	//nameChnl <- "kl"
	//// Close the channel after sending all names to signal completion
	//close(nameChnl)
	//wg.Wait()

	go test1(nameChnl)
	fmt.Println(<-nameChnl)
	fmt.Println(<-nameChnl)

}
func test1(ch chan string) {
	ch <- "papun"
	ch <- "saswat"
}

func printNames(name string, nameChnl chan bool) {
	time.Sleep(time.Second * 1)
	fmt.Println(name)
	nameChnl <- true
}
func processNames(ch <-chan string, wg *sync.WaitGroup) {
	for name := range ch {
		// Simulate processing time (replace with your actual logic)
		processName(name)
	}
	wg.Done()
}

// Simulate processing time (replace with your actual logic)
func processName(name string) {
	// Replace with your actual processing logic
	fmt.Println("start Processing name:", name)
	time.Sleep(time.Second * 2) // Simulate processing time
	fmt.Println("end Processed name:", name)
}
