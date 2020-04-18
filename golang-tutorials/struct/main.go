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
	// Another way is to just use 'contactInfo' which is equivalent to 'contactInfo contactInfo'
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

	jim.print()
	// this assigns the memory address of jim to jimPointer
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	// the below will still update the first name becasue the reciever is a pointer to person
	//jim.updateName(Jimmy)

	jim.print()
}

// Go is "Pass By Value"
// Turn address into value with *address
// Turn value into address with &value

// *type  - this is a type declaration , it means we are working with a pointer to a person
// *variable - this is an operator , it means we want to manipulate the value the pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
