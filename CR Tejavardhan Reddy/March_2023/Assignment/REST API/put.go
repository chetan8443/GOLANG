package main

import (
	"fmt"
	"net/http"
)

func main() {
	link := "http://google.com"
	checkLink(link)
}

func checkLink(link string) {
	fmt.Println("Spiritual")
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
