package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
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

	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)

}
