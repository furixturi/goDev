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
	fmt.Printf("%+v", jim)

}
