package main

import "fmt"

// definition of struct
type person struct {
	firstName string
	lastName  string
	age       int
	addr      address
}

type address struct {
	city string
	zip  int
}

func main() {
	// declartion without init -> initialize with zero values for all props
	// var p person

	// declare and init
	// p := person{"Rajesh", "Reddy", 12}
	// p := person{firstName: "Rajesh", lastName: "Reddy", age: 12}
	p := person{lastName: "Reddy", firstName: "Rajesh"}
	p.age = 12

	p.age = 15
	p.print()

	np := person{
		firstName: "Ramesh",
		lastName:  "Test",
		age:       22,
		addr: address{
			city: "Phoenix",
			zip:  85021,
		},
	}
	fmt.Printf("%+v", np)
}

// method with a reciever type person
// person p is passed by value
func (p person) print() {
	fmt.Printf("%+v", p)
}
