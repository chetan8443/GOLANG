/*
package main

import "fmt"

	func main() {
		const length int = 90
		const breadth int = 70
		var area int = length * breadth
		fmt.Println("The area of the rectangle is", area)
	}
*/
/*package main

import "fmt"

func main() {
	fmt.Println("Enter your First Name:")
	var first string
	fmt.Scanln(&first)
	fmt.Println("Enter your Second Name")
	var second string
	fmt.Scanln(&second)
	fmt.Println("Hi My dear", first+" "+second)
}*/
/*package main

import "fmt"

func main() {
	var a string
	fmt.Println("Enter the String my dear user")
	fmt.Scanln(&a)
	if a == "HarshithaTangala" {
		fmt.Println("Correct Answer!")
	} else if a == "Harshitha" {
		fmt.Println("You are too near Guess fast!")
	} else {
		fmt.Println("Type the correct answer!")
	}
}*/
package main

import "fmt"

func main() {
	var n int = 10
	var i int
	for i = 1; i <= n; i++ {
		fmt.Println(i)
	}
}
