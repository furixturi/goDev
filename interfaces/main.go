package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (eb englishBot) getGreeting() string {
	// VERY custom logic to generate an English greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// VERY custom logic to generate an Spanish greeting
	return "Hola!"
}
