/*
HackerLand University has the following grading policy:
Every student receives a grade in the inclusive range from 0 to 100 .
Any grade less than 40 is a failing grade.
Examples

	grade = 83 round to 85 (85 - 84 is less than 3)
	grade = 29 do not round (result is less than 40)
	grade = 57 do not round (60 - 57 is 3 or higher)

Sample Input :
4
73
67
38
33
sample Output:
75
67
40
33
*/
package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		if t < 38 {
			fmt.Println(t)
		} else {
			if t%5 > 2 {
				// round our number
				fmt.Println(((t / 5) + 1) * 5)
			} else {
				fmt.Println(t)
			}
		}
	}
}
