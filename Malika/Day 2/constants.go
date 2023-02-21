/*
Write a program that declares and initializes constants of different data types
in Golang
*/
package main

import "fmt" // importing fmt

const integer int = 80                // declaring and initializing integer constants
const float float64 = 35.63           // declaring and initializing float constants
const boolean bool = true             // declaring and initializing boolean constants
const String string = "Assignment 2." // declaring and initializing string constants

// main function
func main() {

	fmt.Println("These are constants")
	fmt.Printf("This is an Integer number : %d \n", integer) //prints out integer
	fmt.Printf("This is a float number : %f \n", float)      //prints out float
	fmt.Printf("This is a boolean value : %t \n", boolean)   //prints out boolean
	fmt.Printf("This is a string : %s \n", String)           //prints out string
}
