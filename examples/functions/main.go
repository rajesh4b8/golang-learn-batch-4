package main

import "fmt"

func main() {
	// calling the function with `()`
	simpleFunction()

	withParams("str", 133)
	withParamsOfSameType(12, 133)

	age := retuningASingleValue()
	fmt.Println("Age:", age)

	a, b := retuningMultiValues()
	fmt.Println(a, b)
	a, b = retuningMultiValues2()
	fmt.Println(a, b)
}

// defining a simple funtion with no return types
func simpleFunction() {
	fmt.Println("Simple Function")
}

// multiple parameters
func withParams(p1 string, p2 int) {
	fmt.Println("Inside withParams :: ", p1, p2)
}

// multiple parameters of same type
func withParamsOfSameType(p1, p2 int) {
	fmt.Println("Inside withParamsOfSameType :: ", p1, p2)
}

// age
func retuningASingleValue() int {
	return 43
}

// name and age
func retuningMultiValues() (string, int) {
	return "Rao", 43
}

// name and age
func retuningMultiValues2() (s string, n int) {
	s = "Rao"
	n = 43
	return
}
