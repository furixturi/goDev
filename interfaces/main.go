package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// If you have a getGreeting function that returns a string in this program
// You are also a bot now
func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// Any type in this program which has a function
// called getGreeting and returns a string,
// is now an honorary bot
// who can be used in this function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
