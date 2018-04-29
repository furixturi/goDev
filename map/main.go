package main

import (
	"fmt"
)

func main() {
	colors := make(map[string]string)

	// To add or access a key value pair in a map,
	// Always use square bracket syntax
	// The dot syntax won't work because keys in map are static typed
	colors["white"] = "#ffffff"

	delete(colors, "white")
	fmt.Println(colors)
}
