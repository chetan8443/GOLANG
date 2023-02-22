package main

import "fmt"

func main() {

	// outer for loop
	for i := 1; i <= 3; i++ {

		// inner for loop
		for j := 1; j <= 3; j++ {

			// terminates the inner for loop only
			if i == 2 {
				break
			}

			fmt.Println("i=", i, "j=", j)
		}
	}
}
