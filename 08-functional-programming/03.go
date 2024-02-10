package main

import "fmt"

type Product2 struct {
	ID           int
	Name         string
	Price        float64
	Category     string
	AvailableQty int
	IsOnSale     bool
	// Add more fields as needed
}

func main() {
	//nums := []int{1, 2, 3, 4, 5}
	//result := Map(func(n int) int {
	//	return n * 2
	//}, nums)

	//names := []string{"saswat", "papun", "virat", "kl", "chahal"}
	//result := Map(func(n string) int {
	//	return len(n)
	//}, names)
	//fmt.Println(result)

	products := []Product2{
		{ID: 1, Name: "Laptop", Price: 1000, Category: "Electronics", AvailableQty: 50, IsOnSale: false},
		{ID: 2, Name: "Book", Price: 20, Category: "Education", AvailableQty: 100, IsOnSale: true},
		{ID: 3, Name: "Headphones", Price: 50, Category: "Electronics", AvailableQty: 30, IsOnSale: true},
		{ID: 4, Name: "Backpack", Price: 40, Category: "Fashion", AvailableQty: 80, IsOnSale: false},
		{ID: 5, Name: "Smartphone", Price: 800, Category: "Electronics", AvailableQty: 60, IsOnSale: true},
	}

	productNames := Map(func(product Product2) string {
		return product.Name
	}, products)
	fmt.Println(productNames)

}

func Map[T, U any](f func(T) U, list []T) []U {
	result := make([]U, len(list))
	for i, v := range list {
		result[i] = f(v)
	}
	return result
}
