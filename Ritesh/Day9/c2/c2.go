//Program to print the number of occurences of a character from a word using Goroutines

package main

import (
	"fmt"
	"sync"
)

func countOccurrences(word string, char rune, count *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, c := range word {
		if c == char {
			*count++
		}
	}
}

func main() {
	word := "hello world"
	char := 'l'
	var count int
	var wg sync.WaitGroup

	for _, c := range word {
		if c == char {
			wg.Add(1)
			go countOccurrences(word, char, &count, &wg)
		}
	}

	wg.Wait()
	fmt.Printf("Number of occurrences of %c: %d", char, count)
}
