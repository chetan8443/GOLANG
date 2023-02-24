// Priting pattern for given range

package main

import "fmt"

func main() {
	for i := 'A'; i <= 'G'; i++ {
		for j := 'A'; j <= i; j++ { //using nested for loop
			fmt.Printf("%c ", j)
		}
		fmt.Println()
	}
}
