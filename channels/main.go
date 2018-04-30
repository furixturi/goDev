package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// make a channel of type string to communicate between child go routines to the main go routine
	c := make(chan string)

	for _, link := range links {
		// A function with blocking code will be called
		// We want to use it in a child go routine
		// So we need to pass the channel to it
		go checkLink(link, c) // initiate a brand new go routine with the "go" keyword
	}
	fmt.Println(<-c)
	fmt.Println(<-c) // Now tell the main routine that we're waiting for two messages from the channel
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Send a message to the channel with <-
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep, it's up"
}
