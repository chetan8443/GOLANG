package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// Create a slice of Person structs
	people := []Person{
		{name: "ABC", age: 10},
		{name: "DEF", age: 20},
		{name: "IJK", age: 30},
	}
	fmt.Println("People:", people)

	// Create a map of Person structs
	peopleMap := map[string]Person{
		"ABC": {name: "ABC", age: 10},
		"DEF": {name: "DEF", age: 20},
		"IJK": {name: "IJK", age: 30},
	}
	fmt.Println("People map:", peopleMap)
}
