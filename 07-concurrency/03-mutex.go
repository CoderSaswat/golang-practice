package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex
var names = []string{"saswat"}

func main() {
	now := time.Now()
	urls := []string{"https://google.com", "https://fb.com", "https://go.dev", "https://github.com"}
	for _, url := range urls {
		waitGroup.Add(1)
		go getStatusCode(url)
	}
	waitGroup.Wait()
	fmt.Println(names)
	fmt.Printf("time taken %s", time.Since(now))
}

func getStatusCode(url string) {
	defer waitGroup.Done()
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("internal server error")
	}
	fmt.Printf("url : %s and status: %d\n", url, response.StatusCode)
	//mutex.Lock()
	names = append(names, url)
	//mutex.Unlock()
}
