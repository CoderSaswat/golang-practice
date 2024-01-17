package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("hi")
	//
	//func() {
	//	fmt.Println("i am annonymous")
	//}()
	//
	//func(firstName string, lastName string) {
	//	fmt.Println("i am " + firstName + " " + lastName)
	//}("saswat", "panda")
	//
	//fullName := func(firstName string, lastName string) string {
	//	return "i am " + firstName + " " + lastName
	//}("sas", "panda")
	//fmt.Println(fullName)
	//
	//var myFullName func(firstName string, lastName string) string
	//
	//myFullName = func(firstName string, lastName string) string {
	//	return "i am " + firstName + " " + lastName
	//}
	//
	//println(myFullName("s", "p"))
	//
	//var divideFn func(x int, y int) (q int, r int)
	//
	//divideFn = func(x int, y int) (q int, r int) {
	//	q = x / y
	//	r = x % y
	//	return
	//}
	////divideFn(10, 5)
	//fmt.Println(divideFn(10, 4))
	//
	//q, r := func(x int, y int) (q int, r int) {
	//	q = x / y
	//	r = x % y
	//	return
	//}(15, 4)
	//fmt.Println(q, r)

	//execute(f1)
	//execute(func() {
	//	fmt.Println("anno 01-function")
	//})
	//execute2(sub, 10, 20)

	//var fn func()
	//fn = getFn()
	//fn()
	//var add func(int, int) int
	//add = getAddFn()
	////became like this
	////add := func(a int, b int) int {
	////	return a + b
	////}
	//result := add(10, 20)
	//fmt.Println(result)

	//var getDymanicFuncV func(a int, b int)
	//getDymanicFuncV = getDynamicFn(add)
	////this will be like this
	////getDymanicFuncV = func(a int, b int) {
	////	fmt.Println("operation started at " + time.Now().GoString())
	////	result := add(a, b)
	////	println(result)
	////	fmt.Println("operation completed at " + time.Now().GoString())
	////}
	//getDymanicFuncV(10, 20)

	//n1, n2 := 10, 20
	//operation := 2
	//var choosenOperationFn func(a int, b int) int
	//switch operation {
	//case 1:
	//	choosenOperationFn = add
	//case 2:
	//	choosenOperationFn = sub
	//}
	//result := choosenOperationFn(n1, n2)
	//println(result)

}

func f1() {
	fmt.Println("f1 invoked")
}
func f2() {
	fmt.Println("f2 invoked")
}
func execute(fn func()) {
	fn()
}

func execute2(fn func(int, int) int, a int, b int) {
	result := fn(a, b)
	println(result)
}

func f3(n1 int, n2 int) int {
	return n1 + n2
}

func add(n1 int, n2 int) int {
	return n1 + n2
}

func sub(n1 int, n2 int) int {
	return n1 - n2
}

func getFn() func() {
	return func() {
		fmt.Println("01-function returned")
	}
}

func getAddFn() func(a int, b int) int {
	return func(a int, b int) int {
		return a + b
	}
}

func getDynamicFn(fn func(a int, b int) int) func(a int, b int) {
	return func(a int, b int) {
		fmt.Println("operation started at " + time.Now().GoString())
		result := fn(a, b)
		println(result)
		fmt.Println("operation completed at " + time.Now().GoString())
	}
}
