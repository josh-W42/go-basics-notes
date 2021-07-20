package main

/*
	When making a test file, make sure
	to always name the file like this (ie, name_test)

	Also, it's important to note that
	testing in go is different then
	testing in say, mocha or selenium.

	When testing your code, it's important for
	you as developer to figure out what is important
	...
	once you figure that out, then you can design your tests.

	When making test functions be sure to follow the
	convention used bellow.
*/

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to Ace of Spaces, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}

}
