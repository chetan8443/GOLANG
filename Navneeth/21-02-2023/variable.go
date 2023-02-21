/*Write a program that declares and initializes variables of different data types
in Golang
*/
package main

import "fmt" // importing fmt

var integer int = 30               // declaring and initializing integer variable
var float float64 = 25.67          // declaring and initializing float variable
var boolean bool = true            // declaring and initializing boolean variable
var String string = "Hello World." // declaring and initializing string variable

// main function
func main() {

	fmt.Println("These are variables")
	fmt.Printf("This is an Integer number : %d \n", integer) //prints out integer
	fmt.Printf("This is a float number : %f \n", float)      //prints out float
	fmt.Printf("This is a boolean value : %t \n", boolean)   //prints out boolean
	fmt.Printf("This is a string : %s \n", String)           //prints out string
}
