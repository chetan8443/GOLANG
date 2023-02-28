package main

import "fmt"

// Define a struct named Person
type Person struct {
	name string
	age  int
}

// Define another struct named Employee that embeds Person
type Employee struct {
	Person
	id    int
	title string
}

func main() {
	// Create an Employee instance
	emp := Employee{
		Person: Person{
			name: "John",
			age:  30,
		},
		id:    1001,
		title: "Software Engineer",
	}

	// Access fields from the embedded Person struct
	fmt.Println(emp.name) // Output: John
	fmt.Println(emp.age)  // Output: 30

	// Access fields from the Employee struct
	fmt.Println(emp.id)    // Output: 1001
	fmt.Println(emp.title) // Output: Software Engineer
}
