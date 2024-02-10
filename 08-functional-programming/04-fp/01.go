package main

import (
	_ "fmt"
)

type Product struct {
	ID           int
	Name         string
	Price        float64
	Category     string
	AvailableQty int
	IsOnSale     bool
	// Add more fields as needed
}

func main() {
	//p1 := Product{ID: 1, Name: "Laptop", Price: 1000, Category: "Electronics", AvailableQty: 50, IsOnSale: false}
	//p2 := Product{ID: 2, Name: "Book", Price: 20, Category: "Education", AvailableQty: 100, IsOnSale: true}
	//p3 := Product{ID: 3, Name: "Headphones", Price: 50, Category: "Electronics", AvailableQty: 30, IsOnSale: true}
	//p4 := Product{ID: 4, Name: "Backpack", Price: 40, Category: "Fashion", AvailableQty: 80, IsOnSale: false}
	//p5 := Product{ID: 5, Name: "Smartphone", Price: 800, Category: "Electronics", AvailableQty: 60, IsOnSale: true}
	//products := []Product{p1, p2, p3, p4, p5}

	//productStream := stream.OfSlice(products)
	//productStream.Count()
	//fmt.Println(productStream.Count())

	//mapped := stream.OfSlice(products).Map(func(e types.T) types.R {
	//	product := e.(Product)
	//	return product.ID
	//}).Filter(func(e types.T) bool {
	//	id := e.(int)
	//	return id%2 == 0
	//}).ToSlice()
	//fmt.Println(mapped)

	//var list *arraylist.List
	//list = arraylist.New()
	//list.Size()
	//list.Add(p1, p2, p3, p4, p5)
	//list.Remove(2)
	//list.Sort(func(a, b interface{}) int {
	//	p1 := a.(Product)
	//	p2 := b.(Product)
	//	if p1.Price > p2.Price {
	//		return 1
	//	} else if p1.Price < p2.Price {
	//		return -1
	//	} else {
	//		return 0
	//	}
	//})

	//list.Sort(func(a, b interface{}) int {
	//	p1 := a.(Product)
	//	p2 := b.(Product)
	//	if p1.Price > p2.Price {
	//		return -1
	//	} else if p1.Price < p2.Price {
	//		return 1
	//	} else {
	//		return 0
	//	}
	//})

	//fmt.Println(list)

	//list.Each(func(index int, value interface{}) {
	//	fmt.Printf("index %v index  prodduct %v\n", index, value)
	//})

	//list.Add("virat")
	//list.Add("true")
	//list.Add("kuldeep")

	//mapped := list.Map(func(index int, value interface{}) interface{} {
	//	name := value.(string)
	//	return len(name)
	//})
	//
	//fmt.Println(mapped)
	//filtered := stream.Of(mapped).Filter(func(e types.T) bool {
	//	length := e.(*arraylist.List)
	//	return length%2 == 0
	//}).ToSlice()
	//fmt.Println(filtered)

	//f := stream.Of(list.Iterator()).Filter(func(e types.T) bool {
	//	p := e.(arraylist.Iterator)
	//	fmt.Println(p.First())
	//	return true
	//}).ToSlice()
	//fmt.Println(f)

	//o := optional.Of(1).Map(func(e types.T) types.R {
	//	return e.(int) * 2
	//})

	//h := hashmap.New()
	//h.Put("name", "saswat")
	//h.Put("age", 26)
	////h.ToJSON()
	//fmt.Println(h.ToJSON())

}
