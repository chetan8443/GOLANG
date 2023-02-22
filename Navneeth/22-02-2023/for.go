//nested for loop to print a pyramid of '*'

package main

import "fmt"

//main function
func main() {
	var a int
	fmt.Println("Enter a number.")
	fmt.Scanf("%d", &a) //takes in number to print a pyramid

	//for loop for top half of pyramid
	for i := 1; i <= a; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	//for loop for bottom half of pyramid
	for k := 1; k <= a; k++ {
		for l := k; l <= a; l++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
