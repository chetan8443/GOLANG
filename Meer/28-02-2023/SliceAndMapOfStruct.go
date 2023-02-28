// Program to create a slice of and map of structs
package main

import (
	"fmt"
)

type StructToSlice struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	//converting struct into slice
	slice1 := []StructToSlice{
		{
			Name:    "Meer",
			city:    "Banglore",
			Pincode: 123456,
		},
	}

	for _, j := range slice1 {
		fmt.Println(j.Name, " ", j.city, " ", j.Pincode)
	}

	// converting struct into map
	a1 := StructToSlice{Name: "Dhiraj", city: "Delhi", Pincode: 240005}
	a2 := StructToSlice{"Aniket", "Dehradun", 220008}
	a3 := StructToSlice{Name: "Vishal", city: "Lucknow", Pincode: 107080}

	sample := map[int]StructToSlice{1: a1, 2: a2, 3: a3}
	fmt.Println(sample)
}
