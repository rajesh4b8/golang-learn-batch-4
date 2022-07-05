package main

func main() {
	newDeck := createNewDeck()
	// printDeck(newDeck)
	newDeck.print()

	newDeck.shuffle()

	newDeck.print()

	// type conversion from []string to deck
	// s := []string{"test"}
	// deck(s).print()
}
