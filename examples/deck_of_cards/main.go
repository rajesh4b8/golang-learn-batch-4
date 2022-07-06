package main

func main() {
	newDeck := createNewDeck()
	// printDeck(newDeck)
	// newDeck.print()

	newDeck.shuffle()

	newDeck.writeToFile("myDeck.txt")

	// TODO Assignment
	// newDeckFromFile := readFromFile("myDeck.txt")

	// newDeck.print()

	// deal, remainig := deal(newDeck, 4)
	// fmt.Println("Deal")
	// deal.print()
	// fmt.Println("Remaining")
	// remainig.print()

	// type conversion from []string to deck
	// s := []string{"test"}
	// deck(s).print()
}
