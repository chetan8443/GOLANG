package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {

	person := Person{Name: "abc", Age: 15, Email: "abc@gmail.com"}
	personJSON, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(personJSON))
	var newPerson Person
	err = json.Unmarshal(personJSON, &newPerson)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(newPerson)

}
