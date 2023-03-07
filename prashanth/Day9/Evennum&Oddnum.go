package main

import (
    "fmt"
    "sync"
)

func printEvenNumbers(wg *sync.WaitGroup, ch chan int, n int) {
    defer wg.Done()

    for i := 0; i < n; i++ {
        ch <- 2 * i
    }
}

func printOddNumbers(wg *sync.WaitGroup, ch chan int, n int) {
    defer wg.Done()

    for i := 0; i < n; i++ {
        ch <- 2*i + 1
    }
}

func main() {
    n := 10 // number of even/odd numbers to print

    // create a channel to communicate between goroutines
    ch := make(chan int)

    // create a waitgroup to synchronize goroutines
    var wg sync.WaitGroup
    wg.Add(2)

    // start two goroutines to print even and odd numbers respectively
    go printEvenNumbers(&wg, ch, n)
    go printOddNumbers(&wg, ch, n)

    // read from channel until we have printed enough even and odd numbers
    for i := 0; i < 2*n; i++ {
        fmt.Println(<-ch)
    }

    // wait for goroutines to finish
    wg.Wait()
}
