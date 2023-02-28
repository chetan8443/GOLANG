// 2.Program to create a slice of and map of structs

package main

import "fmt"

// created struct student
type student struct {
	name string
}

func main() {
	s1 := student{name: "john"}
	s2 := student{name: "doe"}
	s3 := student{name: "nav"}

	// created a slice
	slice := [3]student{s1, s2, s3}

	//created a map
	maps := map[int]student{
		1: s1,
		2: s2,
		3: s3,
	}

	fmt.Println(slice)
	fmt.Println(maps)
}
