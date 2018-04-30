package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// os.Stdout implement the Writer interface, res.Body implements the Reader interface
	io.Copy(os.Stdout, res.Body)
}
