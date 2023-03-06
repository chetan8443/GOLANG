package main

import (
    "fmt"
    "sync"
)

func Occurrences(word string, char rune, wg *sync.WaitGroup) {
    defer wg.Done()
    count := 0
    for _, c := range word {
        if c == char {
            count++
        }
    }
    fmt.Printf("The occurences are", count)
}

func main() {
    var wg sync.WaitGroup
    word := "golang is great"
    char := 'g'
    for _, c := range word {
        if c == char {
            wg.Add(1)
            go Occurrences(word, char, &wg)
            break
        }
    }
    wg.Wait()
}
