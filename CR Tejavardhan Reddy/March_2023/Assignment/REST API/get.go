package main

import (
	"fmt"
	"net/http"
	"os"
)

// working with get method
func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(response)
}
