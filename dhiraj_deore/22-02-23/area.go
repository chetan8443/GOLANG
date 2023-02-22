// Write a program to calculate area of different shapes using switch statement :

package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Print("Select the number of shape to calculate area : \n")
		fmt.Println("1. circle\n2. rectangle\n3. triangle\n4. exit ")
		var num float64
		fmt.Scan(&num)
		fmt.Println()

		switch num {
		case 1:
			fmt.Println("Enter the radius of the circle : ")
			var r float64
			fmt.Scan(&r)
			area := 3.14 * r * r
			fmt.Println("The area of circle is : ", area)

		case 2:
			fmt.Println("Enter the length and width of rectangle : ")
			var l float64
			fmt.Scan(&l)
			var w float64
			fmt.Scan(&w)
			area := l * w
			fmt.Println("The area of rectangle is : ", area)

		case 3:
			fmt.Println("Enter the height and base of triangle : ")
			var l float64
			fmt.Scan(&l)
			var w float64
			fmt.Scan(&w)
			area := l * w * 0.5
			fmt.Println("The area of triangle is : ", area)

		case 4:
			fmt.Println("Thank you")
			os.Exit(1)
      
		default:
			fmt.Println("Please enter a valid option !!")
		}
	}

}
