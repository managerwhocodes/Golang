package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

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

	fmt.Println("######################################")

	respThree, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, respThree.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	return 1, nil
}
