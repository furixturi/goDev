package main

import (
	"fmt"
)

func main() {

	cards := newDeck()

	// call a function which returns two values
	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("================")
	remainingCards.print()
}

func newCard() string {
	return "Ace of Spades"
}
