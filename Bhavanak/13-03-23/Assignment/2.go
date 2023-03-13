// Write a program to print the number of occurrences of a character from a word using Goroutines.
// program to print the number of occurrences of a character from a word using Goroutines.
package main

import (
	"fmt"
	"sync"
)

func countOccurrences(word string, char string, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	count := 0
	for i := 0; i < len(word); i++ {
		if string(word[i]) == char {
			count++
		}
	}
	ch <- count
}

func main() {
	var word string
	var char string
	fmt.Println("Enter a word:")
	fmt.Scan(&word)
	fmt.Println("Enter a character to be search:")
	fmt.Scan(&char)

	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < len(word); i++ {
		wg.Add(1)
		go countOccurrences(word[i:i+1], char, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	count := 0
	for n := range ch {
		count += n
	}

	fmt.Printf("The character %s occurs %d times in the word %s\n", char, count, word)
}
