// nested if : largest among three
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Enter three numbers")
	fmt.Scan(&a, &b, &c)
	if a > b {
		if a > c {
			fmt.Printf("%v is greater", a)
		} else {
			fmt.Printf("%v is greater", c)
		}
	} else {
		if b > c {
			fmt.Printf("%v is greater", b)
		} else {
			fmt.Printf("%v is greater", c)
		}

	}
}

/* sample input and output
Input
Enter three numbers 12 26 23
Output
26 is greater */
