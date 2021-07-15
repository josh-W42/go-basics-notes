/*
	At this point we're going
	to try and create an application
	that has a deck of cards

	However, unlike OO programing
	langauges like python, java, or ruby
	In Go we don't use topics like
	classes, inheritance, or polymorphism

	Instead, Go focuses instead on composition.

	Head over
*/

package main

func main() {
	cards := newDeck()
	// cards.print()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
