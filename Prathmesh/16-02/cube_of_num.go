package main

import "fmt"

func main() {
	fmt.Printf("Program to find the volume of a cube \n")
	var side float64 = 6
	var volume float64 = 0
	volume = side * side * side

	// printing the results
	fmt.Println("Dimension of the side : ", side)
	fmt.Println("Therefore, Volume of cube : ", volume)
}
