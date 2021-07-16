package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
	create a new type of 'deck'
	which is a slice of strings

	basically this says that
	there a new type called deck
	and treat that type like a
	slice of strings

*/
type deck []string

/*
	Now we're trying to print
	out all values of the
	deck and to that, we're
	going to make a receiver.

	The receiver is that first bit
	func (variable type) functionName() returnType {}

	The receiver variable is a copy of
	the deck we're working with.

	NOTE: This is how to pass a parameter by value.
	IF you wanted to pass the variable by reference
	you'd probably need an address &variable

	if you recall:
	& - address of value in memory
	* - specifies a pointer

	The receiver type indicates that every
	variable of that type can call this function
	on itself.

	The receiver variable is typically
	one letter or a shortened abbreviation of
	the type.
*/

// creates and returns a new deck
// doesn't need a receiver
// like a constructor, but isn't a constructor
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clovers"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	/*
			when you have a variable that you know
		 	is being returned but you don't want to
			use, you can use "_" and the go compiler
			will know to ignore them.
	*/
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// prints the contents of a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// the bottem return type syntax is how
// you'd configure the function to return
// values

// this function will split the deck by
// a given hand size
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// an important thing to not here is that
// these slices can be segmented into subsets
// using syntax similar to what you'd find in
// python for lists

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	d.toString()
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
