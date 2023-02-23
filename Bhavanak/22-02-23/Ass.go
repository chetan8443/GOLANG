// Assignment 9
// Golang Program to show the Switch case with break and continue in for loop.
package main

import "fmt"

func main() {
forloop:
	for number := 1; number < 10; number++ {
		fmt.Printf("%d", number)
		switch {
		case number == 1:
			fmt.Println("-- One")
		case number == 2:
			fmt.Println("-- Two")
		case number == 3:
			fmt.Println("-- Three")
		case number == 4:
			fmt.Println("-- Four")
		case number == 5:
			fmt.Println("-- Five")
		case number == 6:
			fmt.Println("-- Six")
			continue
		case number == 7:
			fmt.Println("-- Greater than two")
		case number == 8:
			fmt.Println("-- Eight")
			break forloop
		case number == 9:
			fmt.Println("-- Nine")
		default:
			fmt.Println("-- Number not identified")
		}
	}
}
