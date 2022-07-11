package main

import "fmt"

func main() {
	// Map declaration -> it's going to be a nil
	// var shapes map[string]int

	shapes := make(map[string]int)

	shapes["Triangle"] = 3

	colors := map[string]string{
		"Red":  "#143265",
		"Blue": "#643354",
	}
	colors["Red"] = "#85634"
	updateColor(colors)

	fmt.Println(shapes)
	printMap(colors)

	// compilation error -> It has to be same type of map
	// printMap(shapes)
}

func updateColor(c map[string]string) {
	c["Black"] = "#000000"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, ":", hex)
	}
}
