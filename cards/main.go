package main

func main() {
	// to use the new type deck
	// now we need to say
	// go run main.go deck.go
	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
