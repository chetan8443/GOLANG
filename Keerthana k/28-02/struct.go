package main

import "fmt"

type person struct {
	fname string
	lname string
}
type employee struct {
	empid int
	person
}

func main() {
	e1 := employee{1, person{"abc", "xyz"}}
	fmt.Printf("% +v", e1)
}
