package main

// write a program to remove a char from given string
import "fmt"

func main() {
	removeChar()
}

func removeChar() {
	var str string
	fmt.Println("Enter the string")
	fmt.Scan(&str)
	var char string //taking input from user
	fmt.Println("Enter the char to remove from the string")
	fmt.Scan(&char)
	var result string
	var count int = 0
	for _, v := range str {
		if string(v) == char { // checking for perticular char in string
			count++
			continue
		}

		result += string(v)
	}
	if count == 0 {
		fmt.Println("char not found in string")
	}
	fmt.Println(result) //Printing final result
}
