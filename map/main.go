package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// map is a reference type, its keys are indexed so we can loop
// keys can be deleted after added

// struct is a value type, its keys aren't typed nor indexed
// all keys must be known at compile time
