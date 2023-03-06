package main

import (
	"fmt"
	"sync"
)

func countOccurrences(wg *sync.WaitGroup, ch chan int, word string, char rune) {
	defer wg.Done()
	count := 0
	for _, c := range word {
		if c == char {
			count++
		}
	}
	ch <- count
}

func main() {
	word := "hello world"
	char := 'l'

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(word))

	for _, c := range word {
		go countOccurrences(&wg, ch, word, c)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	total := 0
	for count := range ch {
		total += count
	}
	fmt.Printf("The character '%c' appears %d times in '%s'.\n", char, total, word)
}
