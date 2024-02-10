package main

import "fmt"

type Product2 struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	//p1 := Product2{
	//	Id:   0,
	//	Name: "pen",
	//	Cost: 20,
	//}
	//p1.Greet()
	//fmt.Println(p1.PrintProduct())
	//p1.ApplyDiscount(50)
	//fmt.Println(p1.PrintProduct())

}

func (Product2) Greet() {
	fmt.Println("i am a product")
}
func (product Product2) PrintProduct() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", product.Id, product.Name, product.Cost)
}
func (product *Product2) ApplyDiscount(discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}
