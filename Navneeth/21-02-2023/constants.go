/*Write a program that declares and initializes constants of different data types
in Golang
*/
package main

import "fmt" // importing fmt

const i int = 30        // declaring and initializing integer constants
const f float64 = 25.67 // declaring and initializing float constants

// main function
func Constant() {

	const boolean bool = true            // declaring and initializing boolean constants
	const String string = "Hello World." // declaring and initializing string constants
	fmt.Println("These are constants")
	fmt.Printf("This is a boolean value : %t \n", boolean) //prints out boolean
	fmt.Printf("This is a string : %s \n", String)         //prints out string
	p_global()
}

func p_global() {
	fmt.Printf("This is an Integer number : %d \n", i) //prints out integer
	fmt.Printf("This is a float number : %f \n", f)    //prints out float
}
