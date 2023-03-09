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

	person := Person{"Priya", 28, "123 Hyd"}
	encoded, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encoded data:", string(encoded))

	var decoded Person
	err = json.Unmarshal(encoded, &decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded data:", decoded)
}
