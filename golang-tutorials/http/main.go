package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	respOne, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	fmt.Println(respOne)

	bs := make([]byte, 99999)
	// use the read method to get the response body
	respOne.Body.Read(bs)

	fmt.Println(string(bs))
	fmt.Println("=====================================")

	respTwo, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	// Using writer interface
	io.Copy(os.Stdout, respTwo.Body)
}
