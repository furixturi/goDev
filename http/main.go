package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://google.com")

	// Error handling:
	if err != nil {
		fmt.Println("Error: ", err) // Log the error
		os.Exit(1)                  // Exit the program with code other than 0
	}

	fmt.Println(res)
}
