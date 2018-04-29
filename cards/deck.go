package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{} // an empty deck slice

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Return the error from WriteFile
func (d deck) saveToFile(filename string) error {
	// Write to a file with the filename in the 1st parameter
	// With a byte slice in the 2nd parameter
	// If it doesn't exist, create one
	// With the permission code in the 3rd parameter
	return ioutil.WriteFile(filename, []byte(d.toString()), 777)
}
