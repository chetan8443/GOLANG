package main

import (
	"fmt"
	"sync"
)

func countOccurrences(word string, char byte, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	count := 0
	for i := 0; i < len(word); i++ {
		fmt.Print(word[i])
		if word[i] == char {
			count++
		}
	}
	ch <- count
}

func main() {
	var word string
	var char byte
	fmt.Println("Enter a word:")
	fmt.Scanln(&word)
	fmt.Println("Enter a character:")
	fmt.Scanf("%c", &char)

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

	fmt.Printf("The character %c occurs %d times in the word %s\n", char, count, word)
}
