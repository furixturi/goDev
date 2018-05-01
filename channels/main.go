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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	//	go checkLink(<-c, c)
	// }

	// An alternative way of writing the loop:
	// For every message that the channel c sent back
	for l := range c {
		go checkLink(l, c)
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
