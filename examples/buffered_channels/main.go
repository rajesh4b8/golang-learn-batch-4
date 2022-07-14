package main

import "fmt"

func main() {
	// Non buffered channel
	ch := make(chan int)

	go func() {
		ch <- 1
	}()
	fmt.Println("Non buffed channel", <-ch)

	bc := make(chan int, 4)
	bc <- 1
	bc <- 2
	bc <- 3
	fmt.Println("Buffered channel", <-bc, <-bc)

	ch1 := make(chan int)
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)
	}()

	fmt.Println("Read channel that will be closed..")
	for i := range ch1 {
		fmt.Println(i)
	}
}
