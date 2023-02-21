package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter() // calling greeter function inside main function
	result := adder(3, 5)
	fmt.Println("Result is : ", result)
	proresult := proAdder(3, 4, 5, 6, 20, 21, 22) // calling the function inside main function
	fmt.Println("Proresult values is:", proresult)

}
func adder(valOne int, valTwo int) int { //creating a function with two arguments and returning int value
	return valOne + valTwo // returning a int value
}

func greeter() { // creating another function
	fmt.Println("Namstey from golang")
}

func proAdder(values ...int) int { // creating another function with any number of arguments we can pass
	total := 0

	for _, val := range values { // Iterating through the loop
		total += val
	}
	return total // returning a value
}
