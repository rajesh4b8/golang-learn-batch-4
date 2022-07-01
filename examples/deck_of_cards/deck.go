package main

import "fmt"

type deck []string

func createNewDeck() deck {
	newDeck := deck{"Ace of Spades", "Two of Diamonds"}
	newDeck = append(newDeck, "Three of Clubs")

	// TODO: assignment
	// Return a new deck of cards from all suits (Spades, Diamonds, Clubs, Hearts)
	// and numbers (Ace, two, three, four)

	return newDeck
}

func (d deck) print() {
	// fmt.Println(d)
	for i, card := range d {
		fmt.Println(i, card)
	}
}
