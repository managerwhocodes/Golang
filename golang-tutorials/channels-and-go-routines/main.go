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

	// channel implementation
	c := make(chan string)

	for _, link := range links {
		// this will create Go Routine
		go checkLink(link, c)
	}

	// this will wait till main routine gets a message from channel
	// after that main routine with exit
	fmt.Println(<-c)

	// efficient way of receiving messages
	// the counter starts from 1 and not 0 because we have used one print startement above
	for i := 1; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "error occured accessing the link")
		// passing message to channel
		c <- "Some is wrong !"
		return
	}

	fmt.Println(link, "is accessible")
	// passing message to channel
	c <- "Everything is fine !"
}
