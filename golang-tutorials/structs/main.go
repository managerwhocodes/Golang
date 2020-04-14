package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	sam := person{firstName: "Sam", lastName: "Anderson"}
	fmt.Println(sam)

	// default value of struct
	var alex person
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	var john person
	john.firstName = "John"
	john.lastName = "Anderson"
	fmt.Println(john)
	fmt.Printf("%+v", john)
}
