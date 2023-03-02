package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {

		fmt.Println(link, "is might down !")
		c <- "yep it's  might down!"
		return
	}
	fmt.Println(link, "is Working !")
	c <- "yep it's up !"
}
