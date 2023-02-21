package main

import "fmt"

const a = 20         // global constant of int datatype
const b = "Infobell" // global constant of string type
func main() {
	const a = 50.15                                        // local constant of float64 type
	const b = false                                        // constant of
	fmt.Println("value of local variable a and b :", a, b) // printing valus of local connstants
	fmt.Printf("type of local variable a and b :%T, %T \n", a, b)
	printGlobal() // calling function which is outside of the main
}
func printGlobal() { // function will print values and type of global constants
	fmt.Println("global value of variable a and b :", a, b)
	fmt.Printf("type of global variable a and b :%T, %T \n", a, b)
}
