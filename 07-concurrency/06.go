package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Result struct {
	url          string
	responseCode int
}

func main() {
	start := time.Now()
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
		//"https://www.hddibeidghwiddy.com",
	}

	var wg sync.WaitGroup
	var result = make(chan Result)

	for _, url := range urls {
		wg.Add(1)
		go getTitle(url, &wg, result)
	}
	//cllose the channel
	go func() {
		wg.Wait()
		close(result)
	}()
	//wg.Wait()
	//close(result)
	for result := range result {
		//r := <-result
		fmt.Printf("Title of %v: %v\n", result.url, result.responseCode)
	}
	//for {
	//	select {
	//	case r, ok := <-result:
	//		if !ok {
	//			// The channel is closed, exit the loop
	//			fmt.Println("Channel closed. Exiting.")
	//			return
	//		}
	//		fmt.Printf("Title of %v: %v\n", r.url, r.responseCode)
	//	}
	//}
	fmt.Println(time.Since(start))
}

func getTitle(url string, wg *sync.WaitGroup, result chan<- Result) {

	defer wg.Done()
	no, _ := goid()
	response, err := http.Get(url)
	if err != nil {
		r := Result{
			url:          url,
			responseCode: http.StatusNotFound,
		}
		result <- r
		return
	}
	fmt.Println(no)
	r := Result{
		url:          url,
		responseCode: response.StatusCode,
	}
	result <- r
}

var (
	goroutinePrefix = []byte("goroutine ")
	errBadStack     = errors.New("invalid runtime.Stack output")
)

func goid() (int, error) {
	buf := make([]byte, 32)
	n := runtime.Stack(buf, false)
	buf = buf[:n]
	// goroutine 1 [running]: ...

	buf, ok := bytes.CutPrefix(buf, goroutinePrefix)
	if !ok {
		return 0, errBadStack
	}

	i := bytes.IndexByte(buf, ' ')
	if i < 0 {
		return 0, errBadStack
	}

	return strconv.Atoi(string(buf[:i]))
}
