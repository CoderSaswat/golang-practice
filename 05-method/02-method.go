package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type Products []Product
type Predicate func(product Product) bool

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 200, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println(len(products))
	//var expensiveCostFilterPredicate = func(product Product) bool {
	//	return product.Cost > 1000
	//}
	//stationaryProductPredicate := func(product Product) bool {
	//	return product.Category == "Stationary"
	//}
	//fmt.Println(products.Filter(expensiveCostFilterPredicate).format())
	//fmt.Println(products.Filter(stationaryProductPredicate).format())

}

func (products Products) format() string {
	var sb strings.Builder
	for _, p := range products {
		sb.WriteString(fmt.Sprintf("%s\n", p.PrintProduct()))
	}
	return sb.String()
}

func (product Product) PrintProduct() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

//func (products Products) Filter() Products {
//	var result Products
//	for _, product := range products {
//		if product.Cost > 1000 {
//			result = append(result, product)
//		}
//	}
//	return result
//}

//func (products Products) Filter(fn func(product Product) bool) Products {
//	var result Products
//	for _, product := range products {
//		if fn(product) {
//			result = append(result, product)
//		}
//	}
//	return result
//}

func (products Products) Filter(predicate Predicate) Products {
	var result Products
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func getExpensiveProductPredicateInput(cost float32) float32 {
	return cost
}
