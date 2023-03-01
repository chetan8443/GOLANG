package main

import "fmt"

type contact struct {
	email   string
	phoneno int
}

type person struct {
	name        string
	age         int
	contactinfo contact
}

func main() {
	nithin := person{
		name: "Nithin",
		age:  21,
		contactinfo: contact{
			email:   "email@gmail.com",
			phoneno: 3956546545,
		},
	}

	fmt.Printf("Contents of Struct: %+v\n", nithin)
	fmt.Printf("Name of person: %+v\n", nithin.name)
	fmt.Printf("Age of person: %+v\n", nithin.age)
	fmt.Printf("Contact info of person: %+v\n", nithin.contactinfo)
	// fmt.Printf("%+v\n", con)
}