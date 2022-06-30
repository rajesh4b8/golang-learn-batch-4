package main

import "fmt"

var (
	t = 23
	i = 1 // this is defined in global scope and unused
)

func main() {

	// var is the keyword
	// foo is the name of the variable
	// int is the built in type
	// 42 is the value we are initializing with
	var foo int = 42    // declartion with initialization
	var foo1 = 42       // declartion with initialization
	var bar int         // declartion without initalization
	var i, j int = 2, 3 // multi value declartion with initialization
	var x, y int        // multi value declartion without initialization

	// var name string = "Rajesh"
	name := "Rajesh"

	a, b := 2, "test"

	// b is declared as a string, you can't assign an int
	// b = 45

	t = 56 // updating the t, which is defined in global scope

	fmt.Println(foo, foo1)
	fmt.Println(bar)
	fmt.Println(name)
	fmt.Println(i, j, x, y)
	fmt.Println(a, b)
	fmt.Println(t)
}
