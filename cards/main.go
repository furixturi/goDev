package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
	cardsDoesntExist := newDeckFromFile("doesnt_exist")
	cardsDoesntExist.print()
}
