package main

func main() {
	// newDeck := createNewDeck()
	// printDeck(newDeck)
	// newDeck.print()

	// newDeck.shuffle()

	// newDeck.writeToFile("myDeck.txt")

	// TODO Assignment 3
	newDeckFromFile := readFromFile("myDeck.txt")
	newDeckFromFile.print()

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
