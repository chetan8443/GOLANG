package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sentence string
	// sentence = "teja vardhan reddy"
	// upper := "QWERTYUIOPASDFGHJKLZXCVBNM"
	// lower := "qwertyuiopasdfghjklzxcvbnm"
	// specialSymbols:=" ~!@#$%^&*()`-=_+{}[]"
	upperSum := 0
	lowerSum := 0
	symbolSum := 0
	fmt.Printf("Enter the sentence:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		sentence = scanner.Text()
	}
	for _, ch := range sentence {
		//fmt.Println(ch)
		if ch >= 65 && ch <= 90 {
			upperSum += 1
		} else if ch >= 97 && ch <= 122 {
			lowerSum += 1
		} else {
			symbolSum += 1
		}

	}
	fmt.Println("The upper case letters are of ", upperSum, " characters")
	fmt.Println("The lower case letters are of ", lowerSum, " characters")
	fmt.Println("The Special Symbols are of ", symbolSum, " characters")
}
