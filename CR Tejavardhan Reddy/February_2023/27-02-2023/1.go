// A sentence is a string of single-space separated words where each word consists only of lowercase letters.
// A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.
// Given a sentence s return  all the uncommon words. You may return the answer in any order.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sentence string
	//sentence = "teja vardhan reddy"
	fmt.Printf("Enter the string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		sentence = scanner.Text()
	}
	//fmt.Printf("%v\n", sentence)
	res1 := strings.Split(sentence, " ")
	m := map[string]int{}
	for _, word := range res1 {
		m[word] = strings.Count(sentence, word)
	}
	//fmt.Println(m)
	for key, value := range m {
		if value == 1 {
			fmt.Println(key, value)
		}
	}

}
