// program to print the number of occurrences of a character from a word using Goroutines.

package main

import (
	"fmt"
	"sync"
)

func Occurrence(w string, c string, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	count := 0
	for i := 0; i < len(w); i++ {
		if string(w[i]) == c {
			count++
		}
	}
	ch <- count
}

func main() {
	var w string
	var c string
	fmt.Println("Enter word")
	fmt.Scan(&w)
	fmt.Println("Enter character to be searched")
	fmt.Scan(&c)

	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < len(w); i++ {
		wg.Add(1)
		go Occurrence(w[i:i+1], c, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	count := 0
	for n := range ch {
		count += n
	}

	fmt.Printf("The character %s occurs %d times", c, count)
}
