//Creating a pointer to a struct and accessing its fields

package main

import "fmt"

type vehicle struct { // creating a struct with fields
	name    string
	company string
	wheels  int
	fuel    string
}

func main() {
	v1 := vehicle{ 
		"City",
		"Honda",
		4,
		"Petrol",
	}

	fmt.Println(v1)

	changeOwner(&v1) // providing pointer of vehicle struct 
	fmt.Println(v1)

}

func changeOwner(v *vehicle) {
	v.company = "Toyota" // accessing and changing the value of fields of stuct
}
