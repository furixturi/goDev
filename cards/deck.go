package main

import (
	"fmt"
)

// Create a new type of "deck"
// which is essentially a slice of strings
type deck []string

// function to create a new deck
// returns a deck
func newDeck() deck {
	cards := deck{} // an empty deck slice

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// "_" tells go compiler that we know here is a variable but we don't need to use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// return multiple values from one function
func deal(d deck, handSize int) (deck, deck) {
	// aSlice[startIndex:endButNotIncludedIndex]
	// creates a new reference which points to a subsection of a slice
	return d[:handSize], d[handSize:]
}
