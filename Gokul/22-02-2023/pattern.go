//Prints a series upto a count and skips a number(both count and skip entered by the user)
package main

import "fmt"

func main() {
	var count,skip int
	fmt.Println("Enter count")
	fmt.Scan(&count)
	fmt.Println("Enter num to be skipped")
	fmt.Scan(&skip)
	//for loop
	for i := 1; i <= count; i++ {
		if (i == skip) {
			continue
		}
		fmt.Println(i)
	}
}