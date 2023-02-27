//program to print from 1 to 10

package main

import "fmt"

func main() {
	var s string
	fmt.Println("Do you want to print 1 to 10?")
	fmt.Scanf("%s", &s)
	if s == "y" {
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
	} else if s == "n" {
		fmt.Println("Goodbye!!")
	}
}
