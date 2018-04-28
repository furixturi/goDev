package main

import (
	"fmt"
)

func main() {
	// to use the new type deck
	// now we need to say
	// go run main.go deck.go
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
