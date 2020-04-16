package main

import "fmt"

func main() {
	// declare the type of key and value in a map
	colorsOne := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colorsOne)

	// Some other ways to create map
	var colorsTwo map[string]string
	colorsThree := make(map[string]string)

	fmt.Println(colorsTwo)
	fmt.Println(colorsThree)

	colorsThree["white"] = "#ffffff"
	fmt.Println(colorsThree)
}
