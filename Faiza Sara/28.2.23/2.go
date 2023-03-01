package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}

	type object struct {
		name string
	}

	// Slice of Struct
	s := []person{{"Ahana", 21}, {"zara", 23}, {"zoya", 29}}

	// Map of Struct
	m := map[int]object{1: {"phone"}, 2: {"Tab"}, 3: {"Laptop"}}

	fmt.Println("Slice of struct")
	fmt.Println(s)

	fmt.Println("map of struct")

	for i, j := range m {
		fmt.Println(i, j)
	}

}
