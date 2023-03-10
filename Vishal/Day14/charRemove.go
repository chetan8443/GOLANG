package main
// write a program to remove a char from given string
import "fmt"

func main() {
	removeChar()
}

func removeChar() {
	var str = "hello GO"
	var char string //taking input from user
	fmt.Scanln(&char)
	var result string
	var count int=0
	for _, v := range str {
		if string(v)==char {// checking for perticular char in string
			count++
			continue
		}
		if count==0 {
			fmt.Println("char not found in string")
		}
		result+=string(v)
	}
	fmt.Println(result)  //Printing final result
}
