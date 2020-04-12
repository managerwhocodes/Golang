package main

import "fmt"

// This program implements some operations of a card game
// The main functon is where the execution starts
func main() {

	// Declare a Slice
	// Slice is a like an array that can srink or grow
	// ':=' short assignment statement can be used in place of a var declaration with implicit type
	cardsOne := []string{"Ace of Diamonds", newCard()}

	// append data to the slice
	cardsOne = append(cardsOne, "Six of Spades")

	// 'i' is index of element in the slice,
	//'card' is the current value of element in slice,
	//'range' is the keyword used to iterative over the elements in the slice
	for i, card := range cardsOne {
		fmt.Println(i, card) // if we don't use the defined valiables it will say - declared and not used
	}

	// Once we have created a type 'deck' which is a slice of string
	// we can use it as a type
	cardsTwo := deck{"Ace of Diamonds", newCard()}
	cardsTwo.print()

	cardsThree := newDeck()
	// this will return two values
	hand, remainingCards := deal(cardsThree, 5)
	hand.print()
	remainingCards.print()

	fmt.Println(cardsThree.toString())

	cardsFour := newDeck()
	cardsFour.saveToFile("my_cards")

	cardsFive := newDeckFromFile("my_cards")
	cardsFive.print()

	cardsSix := newDeck()
	cardsSix.shuffle()
	cardsSix.print()
}

// function with return type as String
func newCard() string {
	return "Five of Diamonds"
}
