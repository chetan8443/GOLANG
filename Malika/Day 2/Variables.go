package main

//This section tells the go compiler that the
//package should compile as an executable program
//This section tells that the "fmt = format" package is being used in the program

import (
	"fmt"
)

var x int = 20
var y bool = true
var z string = "Learning Golang"
var ab float64 = 776.87

// Declaration and assignment of the variables
func main() {
	fmt.Println(x)  //Printing the assigned value for the variable "x"
	fmt.Println(y)  //Printing the assigned value for the variable "y"
	fmt.Println(z)  //Printing the assigned value for the variable "z"
	fmt.Println(ab) //Printing the assigned value for the variable "ab"

}
