package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// person p created in memory
	p := person{"Rajesh", 12}

	fmt.Printf("Person data :: %+v", p)

	personPointer := &p
	fmt.Printf("\nPerson pointer addr :: %p", personPointer)
	fmt.Println("\np", personPointer)

	valueAtPointer := *personPointer
	fmt.Printf("Person value at addr :: %+v", valueAtPointer)

	p.updateNameByValue("Suresh")
	fmt.Println("\nupdateNameByValue Person", p)

	// pointer indirection -> automatically pass pointer
	// personPointer.updateNameByPointer("Suresh")
	p.updateNameByPointer("Suresh")

	fmt.Println("\nupdateNameByPointer Person", p)
}

// pass by value -> Value receiver -> data is copied to new address
func (p person) updateNameByValue(newName string) {
	fmt.Printf("\nupdateNameByValue pointer addr :: %p", &p)
	p.name = newName
}

// pass by pointer reciver -> the pointer refers to the same data -> data not copied
func (pp *person) updateNameByPointer(newName string) {
	fmt.Printf("\nupdateNameByPointer pointer addr :: %p", pp)
	// (*pp).name = newName
	pp.name = newName
}
