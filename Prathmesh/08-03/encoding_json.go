# Reading json file with 

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
	// Define some sample data
	person := Person{
		Name:    "Prathmesh",
		Age:     23,
		Address: "mumbai",
	}

	// Convert the person struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	// Convert the JSON data back to a person struct
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("New person:", newPerson)
}
