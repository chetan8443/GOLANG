// write a function to print first n prime numbers

package main

import "fmt"

func main() {
	fmt.Print("Enter the number : ")
	n := 0
	fmt.Scan(&n)

	prime1(n)
}

func prime1(n int) {
	i := 0
	num := 2
	for {
		c := 0
		for k := 2; k < num; k++ {
			if num%k == 0 {
				c++
			}
		}
		if c == 0 {
			fmt.Println(num)
			i++
		}
		if i == n {
			break
		}
		num++

	}

}
