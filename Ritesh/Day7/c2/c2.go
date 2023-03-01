//Program to create a slice and map of structs

// Golang program to show how to
// use structs as map keys
package main

// importing required packages
import "fmt"

// declaring a struct
type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	// Creating struct instances
	a1 := Address{
		Name:    "Ritesh",
		city:    "Delhi",
		Pincode: 831005,
	}
	a2 := Address{
		Name:    "Kumar",
		city:    "Manesar",
		Pincode: 120052,
	}
	a3 := Address{
		Name:    "Mohanty",
		city:    "Gurgaon",
		Pincode: 122001,
	}

	// Declaring a map
	var mp map[Address]int

	// Checking if the map is empty or not
	if mp == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	// Declaring and initialising
	// using map literals
	sample := map[Address]int{a1: 1, a2: 2, a3: 3}
	fmt.Println(sample)
}
