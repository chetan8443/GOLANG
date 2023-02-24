package main

//This section tells the go compiler that the
//package should compile as an executable program

import (
	"fmt"
) //This section tells that the "fmt = format" package is being used in the program

var a int = 10
var b bool = true
var c string = "Learning Go"
var d float64 = 7687678765576.8768768

// In this section we are declaring and assigning the variables

func main() {
	fmt.Println(a) //Printing the assigned value for the variable "a"
	fmt.Println(b) //Printing the assigned value for the variable "b"
	fmt.Println(c) //Printing the assigned value for the variable "c"
	fmt.Println(d) //Printing the assigned value for the variable "d"
	cons()         //Calling the function "cons() written in constant.go program"

}
