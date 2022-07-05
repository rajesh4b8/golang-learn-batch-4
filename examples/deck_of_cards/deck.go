package main

import "fmt"

type deck []string

func createNewDeck() deck {
	// newDeck := deck{"Ace of Spades", "Two of Diamonds"}
	// newDeck = append(newDeck, "Three of Clubs")

	// TODO: assignment
	// Return a new deck of cards from all suits (Spades, Diamonds, Clubs, Hearts)
	// and numbers (Ace, two, three, four)
	suits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	numbers := []string{"Ace", "two", "three", "four"}

	newDeck := []string{}
	for _, suit := range suits {
		for _, number := range numbers {
			newDeck = append(newDeck, number+" of "+suit)
		}
	}

	return newDeck
}

// print is method on deck
// d is a receiver type
func (d deck) print() {
	// fmt.Println(d)
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	// TODO Assignment

	// how to switch values?
	d[0], d[15] = d[15], d[0]

	// find a random number between 0 and 15
	// switch for each element
}
