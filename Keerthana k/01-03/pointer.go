package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() { // creating a pointer to a person struct
	p := &person{
		Name: "abc",
		Age:  30,
	}
	fmt.Println("Name:", p.Name) // access the fields of the person struct using pointer
	fmt.Println("Age:", p.Age)
}
