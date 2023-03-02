//Creating a pointer to a struct and accessing its fields

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// create a new Person struct and assign its fields
	p := Person{name: "bunty", age: 55}

	// create a pointer to the Person struct
	ptr := &p

	// access the fields of the Person struct through the pointer
	fmt.Println("Before changing value")
	fmt.Printf("Name: %s, Age: %d\n", ptr.name, ptr.age)

	fmt.Println("After changing value")
	change_val(&p)
	fmt.Println(p)

}

func change_val(v *Person) {
	v.name = "prathmesh"
}
