package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	x := person{age: 21, lastName: "Bhavana", firstName: "Koppula"}
	fmt.Println(x)
	fmt.Println("firstName:", x.firstName)
	fmt.Println("lastName:", x.lastName)
	fmt.Println("age:", x.age)
}
