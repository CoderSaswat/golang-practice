package main

import (
	"fmt"
)

//type Person3 struct {
//	name    string
//	id      int
//	email   string
//	phone   string
//	address Address
//}

type Person3 struct {
	name  string
	id    int
	email string
	phone string
	Address
}

type Address struct {
	id   string
	at   string
	dist string
	pin  string
}

func newPerson(name, email, phone string, idPerson int, at, dist, pin, idAddress string) *Person3 {
	return &Person3{
		name:  name,
		id:    idPerson,
		email: email,
		phone: phone,
		Address: Address{
			at:   at,
			dist: dist,
			pin:  pin,
			id:   idAddress,
		},
	}
}

func main() {
	p := newPerson("saswat", "s@gmail.com", "9853319315", 1, "pkl", "jajpur", "755043", "10")
	//fmt.Println("name:", p.name, "email:", p.email, "phone:", p.phone, "id:", p.id)
	//fmt.Println(p.address.pin)
	fmt.Println(p.dist)
}
