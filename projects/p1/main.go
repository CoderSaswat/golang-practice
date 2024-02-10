package main

import (
	"bytes"
	"fmt"
	_ "math"
	"runtime"
	"strconv"
)

// isPrime checks if a number is prime.
func isPrime(num int) bool {
	//time.Sleep(2 * time.Second)
	fmt.Println("checking isPrime for : ", num)
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// generatePrimes generates prime numbers and sends them to a channel.
func generatePrimes(n1, n2 int, primeChan chan int) {
	//time.Sleep(1 * time.Second)
	fmt.Println("generating prime numbers")
	fmt.Println(getGID())
	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			primeChan <- i
		}
	}
	close(primeChan)
}

// sumPrimes reads prime numbers from a channel and sums them up.
func sumPrimes(primeChan chan int, sumUpChan chan int) {
	//time.Sleep(2 * time.Second)
	fmt.Println("summing up the price")
	fmt.Println(getGID())
	sum := 0
	for prime := range primeChan {
		sum += prime
	}
	sumUpChan <- sum
	close(sumUpChan)
}

func getSquareRoot(sumUpChan chan int, squareRootChan chan int) {
	//time.Sleep(2 * time.Second)
	fmt.Println("getting the square root")
	fmt.Println(getGID())
	result := <-sumUpChan
	squareRootChan <- (result * result)
	close(squareRootChan)
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func main() {
	fmt.Println(getGID())
	n1 := 3
	n2 := 100

	primeChan := make(chan int)
	sumUpChan := make(chan int)
	squareRootChan := make(chan int)

	go generatePrimes(n1, n2, primeChan)
	go sumPrimes(primeChan, sumUpChan)
	go getSquareRoot(sumUpChan, squareRootChan)
	fmt.Println(getGID())

	result := <-squareRootChan
	println(result)
}
