package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Creating a pointer to a Person struct
	p := &Person{Name: "Charlie", Age: 20}

	// Accessing struct fields using the pointer
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)

	// Modifying struct fields using the pointer
	p.Age = 30
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
