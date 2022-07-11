package main

import "fmt"

type person struct {
	name string
}

func main() {

	var j int
	p := &j

	// same as above but more convenient
	// new -> initilize with a zero value and return the address location / pointer
	i := new(int)

	fmt.Println("i", i)
	fmt.Println("p", p)
	fmt.Println(*i)

	map1 := new(map[string]string)
	fmt.Println(map1)
	fmt.Println(*map1)

	// it's panic when you try update a map created using new keyword
	// because the zero value is nil
	// (*map1)["k1"] = "v1"

	map2 := make(map[string]string)
	fmt.Println(map2)
	map2["k1"] = "v1"
	fmt.Println(map2)

	// var p person
	// p1 := &p

	p1 := new(person)
	(*p1).name = "Rajesh"
	fmt.Println(*p1)
}
