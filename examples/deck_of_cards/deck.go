package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

// definition of struct for deck
type deck1 struct {
	suit   string
	number int
}

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

// java world -> static method
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
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

func (d deck) writeToFile(fileName string) {
	str := strings.Join(d, ",")
	err := ioutil.WriteFile(fileName, []byte(str), 0666)
	if err != nil {
		fmt.Println("Error saving the file: ", fileName, err.Error())
		os.Exit(1)
	}
}