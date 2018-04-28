package main

import (
	"fmt"
)

func main() {
	cards := []string{"Ace of Diamonds", newCard()} // A Slice of type string containing elements in {}
	cards = append(cards, "Six of Spades")          // append function doesn't modify the original slice but returns a new slice

	fmt.Println(cards)
}

func newCard() string {
	return "Ace of Spades"
}
