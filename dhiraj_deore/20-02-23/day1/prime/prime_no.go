// write a function to print first n prime numbers

package prime

import "fmt"

func Prime(n int) { // function to list first n prime number
	i := 0
	num := 2
	for {
		c := 0
		for k := 2; k < num; k++ { // checking for divisibility of number
			if num%k == 0 {
				c++
			}
		}
		if c == 0 {
			fmt.Println(num) // printing the prime number 
			i++
		}
		if i == n {
			break // break after reach limit 
		}
		num++

	}

