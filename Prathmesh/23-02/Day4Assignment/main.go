// Find Out How do you call a global variable in other golang files.
package main

import "fmt"

func main() {
	fmt.Println("Its a main function")
	fmt.Println("Calling global variable of demo file in main.go file", Second)
	// calling global variable of demo.go in main function
	demo() // calling demo function of other file
}
