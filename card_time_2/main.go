/*
	In this section we talk about variables
	and a little about functions
*/

package main

import "fmt"

/*
	in the main function, by leaving the space between
 	the parenthese and the brackets blank, we are indicating
	that nothing would be returned
*/
func main() {
	// ok we have 2 ways of writing variables
	var card string = "Ace of Spades"
	cards := "Ace of Spades"
	/*
	 one is shorthand and leaves it up to the compiler
	 one is very descriptive

	 one thing to note is that := is only used for
	 declariation and initialization, but not for
	 reassignment
	*/

	cards = "King of Diamonds"

	card = newCard()

	// here we can also use functions from another
	// file that is associated with this package.
	// This is contray to other langagues where you
	// need to specify the imports.
	fmt.Println(card, cards, getName())
}

/*
	func functionName(arg1: type, ....) returnType {
	}
*/
func newCard() string {
	return "Five of Diamonds"
}
