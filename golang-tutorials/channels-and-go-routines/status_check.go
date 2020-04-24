package main

import (
	"fmt"
	"net/http"
)

func statusCheck() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// channel implementation
	c := make(chan string)

	fmt.Println("Status Check Starts : -----------------")

	for _, link := range links {
		// this will create Go Routine
		go statusCheckLink(link, c)
	}

	// infinte loop
	for {
		go statusCheckLink(<-c, c)
	}
}

func statusCheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "error occured accessing the link")
		// passing message to channel
		c <- link
		return
	}

	fmt.Println(link, "is accessible")
	// passing message to channel
	c <- link
}
