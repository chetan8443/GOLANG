package main

import (
	"fmt"
)

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.om",
	}
	c := make(chan string)

	//for loop  for CheckLinks function calls as much as we have links loop will call CheckLink function
	for _, link := range links {
		go checkLink(link, c)
	}

	//for loop for printing channel message
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	var count int = 6

	for j := 1; j < count; j++ {

		for i := count; i > j; i-- {
			fmt.Print(" ")
		}
		for k := 0; k < j; k++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
	for j := 1; j < count; j++ {

		for k := 0; k < j; k++ {
			fmt.Print(" ")
		}
		for i := count; i > j; i-- {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
