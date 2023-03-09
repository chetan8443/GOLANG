package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type person struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	JerseyNumber int    `json:"jerseyNumber"`
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
	var people []person
	err = json.Unmarshal(contents, &people)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print each person's name and age
	for _, person := range people {
		fmt.Printf("%s is a %s batsmen wearing the jersey number %d .\n", person.Name, person.Type,person.JerseyNumber)
	}
}
