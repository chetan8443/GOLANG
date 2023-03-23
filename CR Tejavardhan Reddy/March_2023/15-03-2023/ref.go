package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter the sentence:")
	var vowels int = 0
	var notvowels int = 0
	var sentence string
	scanner := bufio.NewReader(os.Stdin)
	if scanner.Scan() {
		sentence = scanner.Text()
	}
	//string := strings.Split(" ")
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == 'A' || sentence[i] == 'a' || sentence[i] == 'E' || sentence[i] == 'e' || sentence[i] == 'I' || sentence[i] == 'i' || sentence[i] == 'O' || sentence[i] == 'o' || sentence[i] == 'U' || sentence[i] == 'u' {
			vowels++
		} else {
			notvowels++
		}
	}
	fmt.Println(vowels, notvowels)
}
