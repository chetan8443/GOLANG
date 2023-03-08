package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// Create a Person struct
	person := Person{
		Name:    "Nav",
		Age:     21,
		Address: "Bangalore 1234.",
	}

	// Marshal the struct into a JSON byte slice
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Print the JSON string
	jsonString := string(jsonBytes)
	fmt.Println("JSON string:", jsonString)

	// Unmarshal the JSON byte slice back into a Person struct
	var newPerson Person
	err = json.Unmarshal(jsonBytes, &newPerson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the new Person struct
	fmt.Println("New person:", newPerson)
}
