package main

import "fmt"

func main() {
	// Code to check behaviour of differ statement
	defer fmt.Print("Fist Print Statement\n")
	defer fmt.Println("Commen Print 1")
	defer fmt.Println("Commen Print 2")
	fmt.Println("Second Print Statement")
	deferExa()

	// Code to write and read contents of the file
	content := "My name is Aniket"
	file := "test.txt"
	writeFile(content, file)
	res := readFile(file)
	fmt.Println(res)
}
