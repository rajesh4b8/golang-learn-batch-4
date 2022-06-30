package main

import "fmt"

const (
	DATE_FORMAT = "mm/dd/yyy" // it's exported as starting with upper case
	localConst  = 34          // it's only inside this package as it's starting with lower case
)

func main() {
	const name string = "Rajesh"
	const name1 = "Sri"
	const i = 200

	fmt.Println(name)
	fmt.Println(DATE_FORMAT)
	fmt.Println(localConst, i)
}
