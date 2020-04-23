package main

import "fmt"

func main() {

	// define a slice of numbers
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// check even/odd
	for number := range numbers {
		if number%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
