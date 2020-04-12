package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// If we don't need a variable and just a placeholder then use '_'
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// There is a Receiver attached here
// 'd' is the actual copy of the deck is available. Here it is reference to cards
// Every variable of type 'deck' can call this function on itself
func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}

// this will return two values
func deal(d deck, handSize int) (deck, deck) {
	// splicename[include the index: exclude the index]
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Join concatenates the elements of its first argument to create a single string.
	// The separator string sep is placed between elements in the resulting string.
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// WriteFile writes data to a file named by filename.
	// If the file does not exist, WriteFile creates it with permissions
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// ReadFile reads the file named by filename and returns the contents.
	bs, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error : ", error)
		os.Exit(1)
	}
	// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {

	// NewSource returns a new pseudo-random Source seeded with the given value
	// A Source represents a source of uniformly-distributed pseudo-random int64 values in the range [0, 1<<63).
	source := rand.NewSource(time.Now().UnixNano())
	// New returns a new Rand that uses random values from src to generate other random values.
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// this is a way to swap the values
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
