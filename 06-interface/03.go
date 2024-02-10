package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//myFunc("a")
	//myFunc(45)
	//myFunc(true)

	//var val interface{} = "saswat"
	//var val interface{} = 10
	//var val interface{} = 1.1
	//myFunc(val)

	response, _ := http.Get("http://www.google.com")
	//bs := make([]byte, 99999)
	//response.Body.Read(bs)
	//fmt.Println(string(bs))

	os.File{}

	io.Copy()
}
func myFunc(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println(a, ":this is integer")
	case string:
		fmt.Println(a, ":this is string")
	case bool:
		fmt.Println(a, ":this is boolean")
	case float64:
		fmt.Println(a, "this is float64")

	}
}
