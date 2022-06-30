package main

import "fmt"

func main() {
	// calling the function with `()`
	simpleFunction()

	withParams("str", 133)
}

// defining a simple funtion with no return types
func simpleFunction() {
	fmt.Println("Simple Function")
}

func withParams(p1 string, p2 int) {
	fmt.Println("Inside withParams :: ", p1, p2)
}
