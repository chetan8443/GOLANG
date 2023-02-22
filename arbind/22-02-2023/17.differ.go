package main

import "fmt"

func main() { //creating a main function
	defer fmt.Println("World") // printing a statement using defer
	defer fmt.Println("One")   // printing a statement using defer
	defer fmt.Println("Two")   // printing a statement using defer
	fmt.Println("Hello")       // printing a statement using defer
	myDefer()                  // calling a function inside main function
}

func myDefer() { // creating a function
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
