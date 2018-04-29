package main

import (
	"fmt"
)

// Definition of a new type person which is essentially a struct
type person struct {
	firstName string // Definition: No colon, no comma
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()
	jim.updateName("Alex")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName //This does not work
	// Because go functions pass parameters by value
	// so the p inside updateName function is not the same instance
	// as the reveiver given into in
	// but a copy of its value
}
