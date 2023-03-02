//Creating a pointer to a struct and accessing its fields

package main

import "fmt"

//creating astruct
type student struct {
	name        string
	id          int
	email       string
	phonenumber int
}

func main() {
	//creating instance
	a := student{name: "Bhavana", id: 101, email: "bhvanana@gmail.com", phonenumber: 1234567890}
	//pointer to struct
	ptr := &a
	fmt.Println(ptr)
	//Printing the address of a
	fmt.Println(&ptr)
	//accessing the struct elements
	fmt.Println(ptr.name)
	//dereferencing
	fmt.Println((*ptr).name)

}
