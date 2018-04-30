package main

import (
	"fmt"
)

// Definition of a new type person which is essentially a struct
type person struct {
	firstName   string // Definition: No colon, no comma
	lastName    string
	contactInfo // We made a custom type called contactInfo
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

	jim.updateName("Jimmy") // Go's short hand way of getting the pointer of a value then pass the pointer to the function which expects a pointer of that value type

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// *type is a pointer type of values that points to a value of that specific type
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // * operator gives the in the RAM address
}
