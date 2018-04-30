package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Initiate an empty byte slice of length 99999
	// Because the Read function of the http response
	// body cannot auto shrink or expand
	// the byte slice it is given
	bs := make([]byte, 99999)
	res.Body.Read(bs)
	fmt.Println(string(bs))
}
