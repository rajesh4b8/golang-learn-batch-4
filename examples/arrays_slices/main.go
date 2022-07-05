package main

import "fmt"

func main() {
	// create array
	// Create array and init to zero values
	var ai [10]int

	// create array and init with zero values
	as := [10]string{}

	ai[0] = 1
	ai[4] = 5
	ai[9] = 10
	fmt.Println(ai)
	fmt.Println(as)

	// Throw an error array index 10 out of bounds [0:10]
	// ai[10] = 45

	// arrWithValues := [5]int{1, 2, 3, 4, 5}
	// arrWithValues := [...]int{1, 2, 3, 4, 5, 6}
	arrWithValues := [5]int{0: 1, 4: 5}
	fmt.Println(arrWithValues)

	// You can't do this with arrays
	// arrWithValues = append(arrWithValues, 12)

	// simple slice from array from index 0 to index 3
	slice1 := ai[0:3]

	//  slice from array from index 0 to index 3
	slice1 = ai[:3]

	//  slice from array from index 3 to cap of array
	slice1 = ai[3:]

	//  slice from array from index 0 to cap of array
	slice1 = ai[:]

	fmt.Println("Lenth:", len(slice1), "Cotenents:", slice1, "Capacity:", cap(ai))

	// create slice with make()
	s1 := make([]int, 10)
	fmt.Println("Lenth:", len(s1), "Cotenents:", s1, "Capacity:", cap(s1))
	// add values to slice
	s1 = append(s1, 12)
	fmt.Println("Lenth:", len(s1), "Cotenents:", s1, "Capacity:", cap(s1))

	// creating a slice
	s2 := []int{}
	fmt.Println("Lenth:", len(s2), "Cotenents:", s2, "Capacity:", cap(s2))
	s2 = append(s2, 12)
	fmt.Println("Lenth:", len(s2), "Cotenents:", s2, "Capacity:", cap(s2))

	s3 := []int{1, 4, 6, 86, 4, 8}
	fmt.Println("Lenth:", len(s3), "Cotenents:", s3, "Capacity:", cap(s3))

	sArray := [10]string{"array"}
	fmt.Println(sArray)
	updateArray(sArray) // value not updated
	fmt.Println(sArray)

	slice4 := sArray[:]
	fmt.Println(slice4)
	updateSlice(slice4) // value is updated
	fmt.Println(slice4)
	fmt.Println(sArray)
}

// arrays are value type -> copied to the function
func updateArray(array [10]string) {
	array[0] = "Updated"
}

// slices are ref data type -> copied to a new slice but it refers to same array
func updateSlice(slice []string) {
	slice[0] = "Updated"

}
