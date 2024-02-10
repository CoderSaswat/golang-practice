package main

import (
	"fmt"
	"net"
)

func main() {
	//result := add(1, 2)
	//result := add(1.6, 2.4)
	//result := add("saswat ", "panda")
	//fmt.Println(result)

	list := []string{"sk", "panda", "saswat", "virendra"}

	//res := Map(func(t string) int {
	//	return len(t)
	//}, list)

	r := Filter(func(name string) bool {
		return len(name) > 2
	}, list)
	fmt.Println(r)

	net.Listen()
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func Map[T, U any](f func(T) U, list []T) []U {
	result := make([]U, len(list))
	for i, v := range list {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](predicate func(T) bool, list []T) []T {
	var result []T
	for _, v := range list {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
