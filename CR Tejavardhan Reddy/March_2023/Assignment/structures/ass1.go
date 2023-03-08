// Creating a pointer to a struct and accessing its fields

package main

import "fmt"

type person struct {
	fname  string
	lname  string
	age    int
	gender string
}

func main() {

	p1 := person{fname: "Teja", lname: "Vardhan", age: 22, gender: "Male"}
	ptr := &p1
	fmt.Printf("\n%+v\n", ptr)
	ptr.fname = "Reddy"
	ptr.age = 25
	fmt.Printf("%+v\n", ptr)
}
