package main

import "fmt"

// Definition of a new type person which is essentially a struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	// Another way to instantiate a struct
	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex)
}
