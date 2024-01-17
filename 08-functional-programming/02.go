package main

import (
	"strings"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

func filter(products []Product, fn func(Product) bool) []Product {
	var filteredProducts []Product
	for _, product := range products {
		if fn(product) {
			//dynamically replacing the logic that can be applied over a product
			//fn(product) -> product.Price > 500
			//fn(product) -> product.Category == "Electronics"
			//fn(product) -> strings.Contains(product.Name, "El)
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func filerOneProductBasedOnPrice(product Product, price float64) bool {
	return product.Price > price
}
func filerOneProductBasedOnPriceFn(price float64) func(product Product) bool {
	return func(product Product) bool {
		return product.Price > price
	}
}
func filterBasedOnCategory(product Product, category string) bool {
	return product.Category == category
}
func filterBasedOnSearch(product Product, input string) bool {
	return strings.Contains(product.Name, input)
}

func mapp(products []Product, fn func(Product) any) []any {
	var mappedProducts []any
	for _, product := range products {
		mappField := fn(product)
		mappedProducts = append(mappedProducts, mappField)
	}
	return mappedProducts
}

func main() {
	products := getProducts()
	//filter
	//price := 500.0
	//category := "Electronics"

	//imperative
	//filteredProducts := filteredProductsBasedOnPrice(products, 500.0)
	//filteredProducts := filteredProductsBasedOnCategory(products, "Electronics")
	//filteredProducts := searchProducts(products, "Sm")

	// not working // dynamic not happening
	//filteredProduct := filter(products, filerOneProductBasedOnPrice)

	//1
	//filteredProductByPrice := filter(products, func(product Product) bool {
	//	return product.Price > price
	//})
	//2
	//filteredProductByPrice := filter(products, filerOneProductBasedOnPriceFn(price))
	//fmt.Println(filteredProductByPrice)

	//filteredProductByCategory := filter(products, func(product Product) bool {
	//	return product.Category == category
	//})
	//fmt.Println(filteredProductByCategory)

	//map
	//productIds := mapProductsBasedOnId(products)
	//fmt.Println(productIds)
	//productNames := mapProductsBasedOnNames(products)
	//fmt.Println(productNames)

	//mappedProducts := mapp(products, func(product Product) any {
	//	return product.Name
	//	//return product.Name
	//})
	//fmt.Println(mappedProducts)

}

func mapProductsBasedOnNames(products []Product) interface{} {
	var productNames []string
	for _, product := range products {
		productNames = append(productNames, product.Name)
	}
	return productNames
}

func mapProductsBasedOnId(products []Product) []int {
	var ids []int
	for _, product := range products {
		ids = append(ids, product.ID)
	}
	return ids
}

func filteredProductsBasedOnCategory(products []Product, category string) interface{} {
	var filteredProducts []Product
	for _, product := range products {
		if product.Category == category {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func filteredProductsBasedOnPrice(products []Product, price float64) interface{} {
	var filteredProducts []Product
	for _, product := range products {
		if product.Price > price {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func searchProducts(products []Product, input string) interface{} {
	var filteredProducts []Product
	for _, product := range products {
		if strings.Contains(product.Name, input) {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}
func getProducts() []Product {
	products := []Product{
		{1, "Laptop", 800.0, "Electronics"},
		{2, "Coffee Maker", 50.0, "Appliances"},
		{3, "Smartphone", 600.0, "Electronics"},
		{4, "Headphones", 80.0, "Electronics"},
		{5, "Toaster", 30.0, "Appliances"},
	}
	return products
}
