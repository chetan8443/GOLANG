// write  a program in golang to read a json file.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Open the JSON file
	file, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the contents of the file
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the JSON into a slice of Person structs
	var people []Person
	err = json.Unmarshal(contents, &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print each person's name and age
	for _, person := range people {
		fmt.Printf("%s is %d years old\n", person.Name, person.Age)
	}
}
