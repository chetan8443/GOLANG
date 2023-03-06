/*
Golang code for finding the ascii values for the
array of strings and joining all the strings into single string
by the seperator and finding the single string ascii value
Input:
Enter the size of array:4
Enter the elements:Teja
Enter the elements:Vardhan
Enter the elements:Reddy
Enter the elements:CR

Output:
Teja : [84 101 106 97]
Vardhan : [86 97 114 100 104 97 110]
Reddy : [82 101 100 100 121]
CR : [67 82]
The string is: Teja Vardhan Reddy CR
[84 101 106 97 32 86 97 114 100 104 97 110 32 82 101 100 100 121 32 67 82
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var size int
	fmt.Print("Enter the size of array:")
	fmt.Scanln(&size)
	var str = make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Print("Enter the elements:")
		fmt.Scanln(&str[i]) //reading the array of elements
	}
	for _, value := range str {
		val := []byte(value)
		fmt.Println(value, ":", val)
	}
	str1 := strings.Join(str, " ") //joining the array of strings into single string
	fmt.Println("The string is:", str1)
	fmt.Println([]byte(str1))
	bs := []byte(str1)    //converting string into ascii values
	fmt.Print(string(bs)) //converting the bytes to string
}
