//Go program to read the count of each character in the given string and print the even and odd counts sum

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	fmt.Print("Enter a string:")
	// fmt.Scanln(&str)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	var evenSum int = 0
	var oddSum int
	res1 := strings.Split(str, "")
	dict := map[string]int{}
	for _, word := range res1 {
		dict[word] = strings.Count(str, word)
	}
	fmt.Println("The dictionay of character counts:", dict)
	for _, char := range dict {
		if char%2 == 0 {
			evenSum += char
		} else {
			oddSum += char
		}
	}
	fmt.Println("sum of Odd Count=", oddSum)
	fmt.Print("sum of Even Count=", evenSum)
}

//Input:
// Enter a string:tejavardhanreddycr
//Output:
// Odd Count= 16
// Even Count=2
