package main

import (
	"fmt"
	"unsafe"
)

type emty struct {
}

type CurrencyConverterService struct {
}

func main() {
	// var a emty = emty{}
	a := emty{}
	fmt.Println(a)

	i := 12
	name := "Rajesh"

	// Defining and declaring a struct but definition is not reusable
	p := struct {
		name string
		age  int
	}{}
	p.name = "Rajesh"
	p.age = 12

	q := struct {
		name string
		age  int
	}{}
	q.name = "Rajesh"
	q.age = 12

	fmt.Println("person", p)

	fmt.Println("Size of i", unsafe.Sizeof(i))
	fmt.Println("Size of name", unsafe.Sizeof(name))
	fmt.Println("Size of empty struct", unsafe.Sizeof(a))
	fmt.Println("Size of empty struct", unsafe.Sizeof(struct{}{}))
	fmt.Println("Size of bool", unsafe.Sizeof(true))

	x := struct{}{}
	y := struct{}{}
	fmt.Println("x == y", x == y)
	fmt.Println("&x == &y", &x == &y)
	fmt.Println("p == q", p == q)
	fmt.Println("&p == &q", &p == &q)

	fmt.Printf("x at %p, y at %p \n", &x, &y)
	fmt.Printf("p at %p, q at %p \n", &p, &q)

	// one pattern with services
	s := CurrencyConverterService{}
	s.ConverCurrency()
}

func (_ CurrencyConverterService) ConverCurrency() {
	fmt.Println("inside service")
}
