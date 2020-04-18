package main

import "fmt"

// the below example shows the problem with structs and need for an interface
// the function that essentially does the same thing for different structs can't be reused

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

// this function is used by both the stucts
/*func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}
*/

// Same function name with different input parameters is not allowed until it is used with interface
/*func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}*/

func (eb englishBot) getGreeting() string {
	return "Hi There !"
}

func (sp spanishBot) getGreeting() string {
	return "Hola !"
}

// declaring the function that is used by both the structs inside an interface
// this makes englishBot and spanishBot a member of type bot
type bot interface {
	getGreeting() string
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
