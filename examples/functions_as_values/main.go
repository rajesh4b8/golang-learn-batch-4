package main

import "fmt"

func main() {
	printSomething := func() {
		fmt.Println("Print Something!!")
	}
	printSomething()

	add := func(a, b int) int {
		return a + b
	}

	fmt.Printf("%v + %v = %v", 2, 3, add(2, 3))
}

// exactly same as defined above
// func printSomething() {
// 	fmt.Println("Something!!")
// }
