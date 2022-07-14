package main

import (
	"fmt"
	"time"
)

type Vehicle interface {
	getWheels() int
}

type car struct {
	wheels int
}

func (c car) getWheels() int {
	return c.wheels
}

func main() {
	demoSelect()
}

func demoSwitch() {
	var v Vehicle = car{4}
	fmt.Println("Vehicle from struct car", v)

	var a interface{}
	a = 5

	// Switch will always executes from top and only runs the first matching code
	switch a.(type) {
	case int:
		fmt.Println("It's an int")
	case int32:
		fmt.Println("It's an int32")
	default:
		fmt.Println("Not sure!")
	}
}

func demoSelect() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
