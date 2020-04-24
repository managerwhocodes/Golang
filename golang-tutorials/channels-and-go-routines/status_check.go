package main

import (
	"fmt"
	"net/http"
	"time"
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
	/*
		for {
			go statusCheckLink(<-c, c)
		}
	*/

	// alternate way of loop
	for l := range c {
		// don't add sleep here , doing so will block the main routine for this time
		// use function literal
		go func(link string) {
			time.Sleep(2 * time.Second)
			statusCheckLink(link, c)
		}(l) // this is used to invoke the function literal
	}

}

func statusCheckLink(link string, c chan string) {
	// adding sleep here is also not appropriate becasue this function is just use to check the link and return
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
