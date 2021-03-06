package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) { // Define the function literal
			time.Sleep(5 * time.Second)
			checkLink(link, c) //Never try to access variables in other go routines directly, but pass them by value using go's pass by value function feature!
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Now we use the channel to pass the link which was just checked
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// Now we use the channel to pass the link which was just checked
	c <- link
}
