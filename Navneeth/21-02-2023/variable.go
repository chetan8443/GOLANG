/*Write a program that declares and initializes variables of different data types
in Golang
*/
package main

import "fmt" // importing fmt

var integer int = 30               // declaring and initializing integer variable
var float float64 = 25.67          // declaring and initializing float variable

// main function
func main() {

	b:= true            // declaring and initializing boolean variable
	s:= "Hello World." // declaring and initializing string variable
	fmt.Println("These are variables")
	fmt.Printf("This is an Integer number : %d \n", integer) //prints out integer
	fmt.Printf("This is a float number : %f \n", float)      //prints out float
	fmt.Printf("This is a boolean value : %t \n", b)   //prints out boolean
	fmt.Printf("This is a string : %s \n", s)           //prints out string
}
