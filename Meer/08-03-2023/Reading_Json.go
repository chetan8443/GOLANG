// Reading JSON data from test.json file
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read the JSON data from file
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))

}
