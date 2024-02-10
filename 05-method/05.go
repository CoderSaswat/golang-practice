package main

import "fmt"

type contactInfo struct {
	email string
	pin   string
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	saswat := person{
		firstName: "saswat",
		lastName:  "panda",
		contactInfo: contactInfo{
			email: "saswat@gamil.com",
			pin:   "755043",
		},
	}
	//personPtr := &saswat
	//personPtr.updateLastName("ptr")
	//saswat.updateLastName("das")
	saswat.updateFirstName("sk")
	fmt.Println(saswat)
}

// actually updating since it is passing the real address
func (personPtr *person) updateLastName(newLastName string) {
	(*personPtr).lastName = newLastName
}

// copy so it will not update
func (person person) updateFirstName(newFirstName string) {
	person.firstName = newFirstName
}
