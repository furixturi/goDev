package main

import (
	"fmt"
)

// Create a new type of "deck"
// which is essentially a slice of strings
type deck []string

// Define a function in the custom type
// whose "receiver" is of the type "deck"
// d is the copy of that type who invokes this function on itself
// it is like "this" in JavaScript
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
