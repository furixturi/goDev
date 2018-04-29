package main

// Definition of a new type person which is essentially a struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	// One way to instantiate a struct
	// Relying on the order to decide the value of each entry
	alex := person{"Alex", "Anderson"}
}
