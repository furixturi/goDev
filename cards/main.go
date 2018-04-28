package main

import (
	"fmt"
)

func main() {
	// initialize a variable and explicitly declare its type,
	// then assign a value to it
	// usable within and outside of functions
	// var card = "Ace of Spades"

	// := initializes and assigns the value of a new variable
	// letting the compiler infer its type by the assigned value
	// be careful not to use := twice on the same variable name
	// otherwise there will be a compiler error
	// Warning: only usable within a function
	card := "Ace of Spades"

	fmt.Println(card)
}
