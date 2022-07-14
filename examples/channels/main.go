package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	start = time.Now()
)

func main() {
	links := []string{
		"http://linkedin.com",
		"http://yahoo.com",
		"http://learningnet.com",
		"http://facebook.com",
		"http://google.com",
	}

	c := make(chan string) // non buffered channel -> no capacity to store the data

	for _, link := range links {
		go checkHealth(link, c)
	}

	// Reading from channel is a blocking operation and the thead sleeps until we get a response
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	for range links {
		fmt.Println(<-c)
	}

	fmt.Println("all done - ", time.Since(start))
}

func checkHealth(link string, c chan string) {
	// Blocking operation -> Go Routines sits idle until it gets a response
	_, err := http.Get(link)
	if err != nil {
		c <- fmt.Sprintf("%v might be down - in %v \n", link, time.Since(start))
		// c <- link + " might be down!"
		return
	}

	// c <- link + " might is up!"

	c <- fmt.Sprintf("%v is up - in %v\n", link, time.Since(start))
}
