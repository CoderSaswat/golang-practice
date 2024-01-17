package calculator

func init() {
	println("calculator initialized - [devide.go]")
}

func Divide(x, y int) int {
	opCount["divide"]++
	println(opCount)
	return x / y
}
