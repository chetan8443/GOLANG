/*
Write a program that declares and initializes constants of different data types
in Golang, including:
- A constant of type `int`
- A constant of type `float64`
- A constant of type `bool`
- A constant of type `string`
*/

package main

import "fmt"

const a = 10         // global constant of int datatype
const b = "Infobell" // global constant of string type

func main() {
	const a = 20.12 // local constant of float64 type
	const b = false // constant of

	fmt.Println("value of local variable a and b :", a, b) // printing valus of local connstants
	fmt.Printf("type of local variable a and b :%T, %T \n", a, b)
	printGlobal()

}

func printGlobal() { // function will print values and type of global constants
	fmt.Println("global value of variable a and b :", a, b)
	fmt.Printf("type of global variable a and b :%T, %T \n", a, b)
}
