// Go program to create an embedded struct
package main

import (
	"fmt"
)

type contactinfo struct {
	email       string
	phonenumber int
}

type person struct {
	firstname string
	lastname  string
	contact   contactinfo
}

/*
func main() {
	a := person{firstname: "Bhavana", lastname: "Hemanth"}
	fmt.Println(a)
	person := contactinfo{email: "bhavana@gmail.com", phonenumber: 1234567890}
	fmt.Println(person)
}
*/

func main() {
	a := person{
		firstname: "Bhavana",
		lastname:  "Hemanth",
		contact: contactinfo{
			email:       "bhavana@gmail.com",
			phonenumber: 1234567890,
		},
	}
	fmt.Printf("%+v", a)
}
