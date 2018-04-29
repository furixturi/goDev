package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 777)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // ReadFile returns a byteslice and an error (nil if no error occured)
	if err != nil {
		// Option 1: log the error and return a call to newDeck() as a fail save solution
		// Option 2: Log the error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	// convert the byte slice to string,
	// then split the string with "," to get a string slice
	return deck(s) // convert the sting slice to a deck, which is essentially the same thing
}
