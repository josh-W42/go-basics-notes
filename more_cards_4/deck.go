package main

import "fmt"

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
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
