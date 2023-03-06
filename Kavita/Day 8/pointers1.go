package main

import "fmt"

type employee struct {
	name string
	ID   int
}

func main() { // creating a  pointer c for struct employee
	c := &employee{name: "harry", ID: 10}

	fmt.Println(c.name, c.ID)

	c.name = "potter" // chnaging the elements in the struct and printing again
	c.ID = 20

	fmt.Println(c.name, c.ID)

}
