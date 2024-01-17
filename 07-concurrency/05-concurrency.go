package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  int
}

var name string = "saswat"

func main() {
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//wg.Add(1)
	//go helloworld(&wg)
	//go goodbye(&wg)
	////time.Sleep(1 * time.Second)
	//wg.Wait()

	//msg := make(chan string)
	//go greet(msg)
	//time.Sleep(10 * time.Second)
	//greeting := <-msg
	//time.Sleep(3 * time.Second)
	//fmt.Println("Greeting received")
	//fmt.Println(greeting)

	//go greet(msg)
	//time.Sleep(time.Second * 2)
	//greeting := <-msg
	//time.Sleep(3 * time.Second)
	//fmt.Println("Greeting received")
	//fmt.Println(greeting)

	//_, ok := <-msg
	//if ok {
	//	fmt.Println("channel is open")
	//} else {
	//	fmt.Println("channel is closed")
	//}

	//c1 := make(chan string)
	//c2 := make(chan string)
	//c3 := make(chan string)
	//go func() {
	//	for {
	//		fmt.Println("before sending f1")
	//		time.Sleep(time.Millisecond * 500)
	//		c1 <- "from 1"
	//		fmt.Println("after sending f1")
	//	}
	//}()
	//go func() {
	//	for {
	//		fmt.Println("before sending f2")
	//		time.Sleep(time.Second * 1)
	//		c2 <- "from 2"
	//		fmt.Println("after sending f2")
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		fmt.Println("before sending f3")
	//		time.Sleep(time.Millisecond * 1500)
	//		c2 <- "from 3"
	//		fmt.Println("after sending f3")
	//	}
	//}()
	//go func() {
	//	for {
	//		select {
	//		case msg1 := <-c1:
	//			fmt.Println("===================" + msg1)
	//		case msg2 := <-c2:
	//			fmt.Println("======================" + msg2)
	//		case msg3 := <-c3:
	//			fmt.Println("===========================" + msg3)
	//		}
	//	}
	//}()

	//c := boring("boring!")
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("You say: %q\n", <-c)
	//}
	//fmt.Println("You're boring; I'm leaving.")

	//joe := boring("Joe")
	//ann := boring("Ann")
	//for i := 0; i < 5; i++ {
	//	fmt.Println(<-joe)
	//	fmt.Println(<-ann)
	//}
	//fmt.Println("You're both boring; I'm leaving.")

	//stringChannel := make(chan string)
	//intChannel := make(chan int)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	stringChannel <- "Hello, Golang!"
	//}()
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	intChannel <- 42
	//}()
	//select {
	//case str := <-stringChannel:
	//	fmt.Println("Received string:", str)
	//case num := <-intChannel:
	//	fmt.Println("Received integer:", num)
	//case <-time.After(5 * time.Second):
	//fmt.Println("Timeout: No data received within 5 seconds.")
	//}
	//time.Sleep(time.Second * 10)

	//ch := make(chan int)
	//go f(ch, 2)
	//for v := range ch {
	//	fmt.Println(v)
	//}

	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//select {
	//case m1 := <-ch:
	//	fmt.Println(m1)
	//}

	//p := Person{"John", 23}
	//ch := make(chan string)
	//go send(ch)
	//for i := 0; ; i++ {
	//	fmt.Println(<-ch)
	//}
	//for value := range ch {
	//	fmt.Println(value)
	//}
	//go SendPerson(ch, p)
	//name := (<-ch)
	//fmt.Println(name)

	channels := make(chan string) // unbuffered channel
	defer close(channels)
	go starter(channels)   // takes in the channel and passes the string
	receiver := <-channels // receiver takes the value in the channel
	follow(receiver)

}

func starter(ch chan string) {
	fmt.Println("This is the starter on call")
	ch <- "Hello,"
}

func follow(starter string) {
	fmt.Println(starter, "From the starter function, This is the follower on call")
}

func send(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- strconv.Itoa(i)
	}
	//close(ch)
}

func SendPerson(ch chan Person, p Person) {
	ch <- p
}
func f(ch chan int, v int) {
	fmt.Println("sending 1")
	ch <- v
	fmt.Println("sending 2")
	ch <- v * 2
	fmt.Println("sending 3")
	ch <- v * 3
	fmt.Println("sending 4")
	ch <- v * 7
	close(ch)
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			fmt.Println("sending")
			c <- fmt.Sprintf("%s %d", msg, i)
			fmt.Println("sent")
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func helloworld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World!")
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")
	ch <- "Hello Rwitesh"
	//close(ch)
	fmt.Println("Greeter completed")
}

func goodbye(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Good Bye!")
}
