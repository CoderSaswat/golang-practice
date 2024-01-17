package main

import "errors"

var ErrDivideByZero error = errors.New("can't be divided by zero")
var ErrDivisorGreaterThanMultiplier = errors.New("can't be divided because divisor is greater than multiplier")

func main() {
	//q, r, err := divide(10, 0)
	//if err != nil {
	//	if errors.Is(err, ErrDivideByZero) {
	//		println(err.Error())
	//		println("changing the flow since ErrDivideByZero occurred")
	//	}
	//	if errors.Is(err, ErrDivisorGreaterThanMultiplier) {
	//		println(err.Error())
	//		println("changfing the flow since ErrDivisorGreaterThanMultiplier occurred")
	//	}
	//}
	//println(q, r)

	//q, r, err := divideService(10, 0)
	//if err != nil {
	//	//if errors.Is(err, ErrDivideByZero) {
	//	//	println(err.Error())
	//	//	println("changing the flow since ErrDivideByZero occurred")
	//	//}
	//	if err == ErrDivideByZero {
	//		println(err.Error())
	//		println("changing the flow since ErrDivideByZero occurred")
	//	}
	//	if errors.Is(err, ErrDivisorGreaterThanMultiplier) {
	//		println(err.Error())
	//		println("changing the flow since ErrDivisorGreaterThanMultiplier occurred")
	//	}
	//}
	//println(q, r)

}

func divideService(a int, b int) (q int, r int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	q, r = divideClient(a, b)
	return
}

func divide(a int, b int) (q int, r int, err error) {
	if b == 0 {
		err = ErrDivideByZero
		return
	}
	if b > a {
		err = ErrDivisorGreaterThanMultiplier
		return
	}
	q = a / b
	r = a % b
	return
}

// 3rd party api
func divideClient(a int, b int) (q int, r int) {
	if b == 0 {
		panic(ErrDivideByZero)
	}
	if b > a {
		panic(ErrDivisorGreaterThanMultiplier)
	}
	q = a / b
	r = a % b
	return
}
