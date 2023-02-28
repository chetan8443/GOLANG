// 1.Program to create an embedded struct

package main

import "fmt"

type person struct { // create a struct
	fname string
	lname string
}

type student struct { //create a embedded struct
	rollno int
	person // using struct as a field
}

func main() {

	s1 := student{1, person{"Alex", "carey"}}

	fmt.Printf("%+v", s1)
}
