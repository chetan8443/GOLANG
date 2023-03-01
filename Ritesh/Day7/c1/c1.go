//To create an embedded struct

package main

import "fmt"

type trainee struct { //creating structure trainee
	fname string
	lname string
}

type programming struct { //creating structure programming
	trainee
	prgknw string
}

func main() {
	te1 := programming{ // Assigning values to the structure
		trainee: trainee{
			fname: "Ritesh",
			lname: "Mohanty",
		},
		prgknw: "No prior programming knowledge",
	}

	fmt.Println(te1) //Printing structure values
}
