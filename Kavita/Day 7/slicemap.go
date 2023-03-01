package main

import "fmt"

type colours struct {
	orange string
	yellow string
}

func main() { //creating the slice from the struct
	c := []colours{
		{orange: "yes"},
		{yellow: "no"},
	}

	newmap := map[int]colours{
		1: {orange: "yes", yellow: "no"},
		2: {orange: "no", yellow: "yes"},
	}

	fmt.Println("slice is:", c)
	fmt.Println("map is ", newmap)
}
