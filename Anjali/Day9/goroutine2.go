package main

import (
	"fmt"
	"sync"
)

func countOccurrences(wg *sync.WaitGroup, word string, char byte) {
	defer wg.Done()
	count := 0
	for i := 0; i < len(word); i++ {
		if word[i] == char {
			count++
		}
	}
	fmt.Printf("Number of occurrences of %c: %d\n", char, count)
}

func main() {
	var wg sync.WaitGroup
	word := "hello"
	for i := 0; i < len(word); i++ {
		wg.Add(1)
		go countOccurrences(&wg, word, word[i])
	}
	wg.Wait()
}
