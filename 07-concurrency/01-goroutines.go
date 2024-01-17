package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	now := time.Now()
	urls := []string{"https://google.com", "https://fb.com", "https://go.dev", "https://github.com"}
	for _, url := range urls {
		getStatusCode(url)
	}
	fmt.Printf("time taken %s", time.Since(now))
}

func getStatusCode(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("internal server error")
	}
	fmt.Printf("url : %s and status: %d\n", url, response.StatusCode)
}
