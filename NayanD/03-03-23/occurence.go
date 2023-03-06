// Write a program to print the number of occurrences of a character from a word using Goroutines.
package main

import (
	"fmt"
	"sync"
)

func occurrences(word string, char byte, g *sync.WaitGroup, ch chan int) {
	defer g.Done()
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

	var g sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < len(word); i++ {
		g.Add(1)
		go occurrences(word[i:i+1], char, &g, ch)
	}

	go func() {
		g.Wait()
		close(ch)
	}()

	count := 0
	for n := range ch {
		count += n
	}

	fmt.Printf("The character %c occurs %d times in the word %s\n", char, count, word)
}
