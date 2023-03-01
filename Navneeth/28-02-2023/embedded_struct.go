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
	navneeth := person{
		name: "Navneeth",
		age:  21,
		contactinfo: contact{
			email:   "email@gmail.com",
			phoneno: 334455662,
		},
	}

	fmt.Printf("Contents of Struct: %+v\n", navneeth)
	fmt.Printf("Name of person: %+v\n", navneeth.name)
	fmt.Printf("Age of person: %+v\n", navneeth.age)
	fmt.Printf("Contact info of person: %+v\n", navneeth.contactinfo)
	// fmt.Printf("%+v\n", con)
}
