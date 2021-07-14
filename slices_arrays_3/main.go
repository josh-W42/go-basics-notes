package main

import "fmt"

func main() {
	/*
		Notes about arrays
		and Slices.

		Arrays - Fixed length lists.
			Has length and capacity.
		Slice - An array that can grow or shrink.

		Bellow is how we create a slice and add values
		to the slice.
	*/

	cards := []string{"Kind of Hearts", newCard()}
	cards = append(cards, "Six of Spades")

	/*
		The for loop:
		i - index
		card - element at array[i]
		range array - the slice of the array
	*/
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
