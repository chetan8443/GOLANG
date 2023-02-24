package main

import "fmt"

func main() {

	// Example 3: Looping until a condition is met
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Example 4: Looping indefinitely with a break statement
	for {
		fmt.Println("Looping indefinitely...")
		break
	}
}
