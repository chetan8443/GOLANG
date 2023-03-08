package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read the JSON data from file
	data, err := ioutil.ReadFile("demo.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("%T", data)

}
