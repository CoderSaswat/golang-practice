package main

import (
	"demo/contract"
	"demo/service"
	"fmt"
)

func main() {
	//fmt.Println("hi")
	var productService contract.ProductService
	productService = service.ProductServiceImpl{}

	fmt.Println(productService.GetProductByID(1))
}
