package main

import "fmt"

type Income interface {
	source() string
	calculate() int
}

type FixedBilling struct {
	projectName string
	bidAmount   int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName       string
	costPerClick int
	noOfClicks   int
}

func (a Advertisement) calculate() int {
	return a.noOfClicks * a.costPerClick
}

func (a Advertisement) source() string {
	return a.adName
}

func main() {
	var p1 Income = FixedBilling{"Project 1", 2000}
	p2 := TimeAndMaterial{"Project 2", 10, 30}
	a1 := Advertisement{"Add1", 2, 13532}

	incomeStreams := []Income{p1, p2, a1}
	calculateNetIncome(incomeStreams)
}

func (p FixedBilling) source() string {
	return p.projectName
}

func (p FixedBilling) calculate() int {
	return p.bidAmount
}

// Method over loading is not supported in go
// func (p FixedBilling) calculate(str string) int {
// 	return p.bidAmount
// }

func (p TimeAndMaterial) source() string {
	return p.projectName
}

func (p TimeAndMaterial) calculate() int {
	return p.hourlyRate * p.noOfHours
}

func calculateNetIncome(streams []Income) {
	var netIncome int = 0
	for _, p := range streams {
		fmt.Printf("Income from %v is %v\n", p.source(), p.calculate())
		netIncome += p.calculate()
	}

	fmt.Println("The Net Income is", netIncome)
}
