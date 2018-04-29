package main

import (
	"fmt"
)

// Definition of a new type person which is essentially a struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	// A third way to instantiate a struct
	var alex person // Declare a variable of type person
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
