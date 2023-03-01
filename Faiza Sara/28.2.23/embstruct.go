// Program to create an embedded struct
package main

import "fmt"

type person struct {
	name  string
	age   int
	alive bool
}

type citizen struct {
	person
	citizenship bool
}

func main() {

	a := person{
		name:  "Anisha",
		age:   22,
		alive: true,
	}

	fmt.Println(a)

	//Embedded Struct

	b := citizen{
		person: person{
			name:  "Reena",
			age:   2,
			alive: true,
		},
		citizenship: false,
	}

	fmt.Print(b)
}
