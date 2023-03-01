// Creating a pointer to a structure and accessing its fields

package main

import "fmt"

type Employee struct { //creating a structure "employee"

	name  string
	empid int
}

func main() {

	emp := Employee{"Ritesh", 1}

	pts := &emp // this is the pointer "pts" to the struct "emp"

	fmt.Println(pts)
	fmt.Println(pts.name)
	fmt.Println((*pts).name)

}
