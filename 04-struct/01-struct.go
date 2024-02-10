package main

import "fmt"

type Person struct {
	Id   string
	Name string
}

type Person2 struct {
	Id   string
	Name *string
}

type Organization struct {
	Id   int
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	p1 := Person{"1", "saswat"}
	p2 := Person{"2", "papun"}
	fmt.Println("Initial p1:", p1.Id, p1.Name) // Output: Initial p1: 1 saswat
	fmt.Println("Initial p2:", p2.Id, p2.Name)
	p2 = p1
	p2.Name = "new name"
	fmt.Println("Modified p1:", p1.Id, p1.Name) // Output: Modified p1: 1 new name
	fmt.Println("Modified p2:", p2.Id, p2.Name)

	name2 := "sas"
	name3 := "pap"

	p3 := Person2{
		Id:   "1",
		Name: &name2,
	}
	p4 := Person2{
		Id:   "2",
		Name: &name3,
	}

	fmt.Println("Initial p3:", p3.Id, *p3.Name)
	fmt.Println("Initial p4:", p4.Id, *p4.Name)
	p3 = p4
	*p3.Name = "will see"
	fmt.Println("Modified p3:", p3.Id, *p3.Name)
	fmt.Println("Modified p4:", p4.Id, *p4.Name)

	org1 := &Organization{
		Id:   10,
		Name: "thoughtclan",
		City: "bangalore",
	}
	emp1 := Employee{
		Id:   1,
		Name: "saswat",
		Org:  org1,
	}
	//fmt.Printf("%+v", emp1)
	fmt.Printf("%+v\n", emp1)

}
