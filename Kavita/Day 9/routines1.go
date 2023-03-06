package main

import (
    "fmt"
    "sync"
)

func Even(wg *sync.WaitGroup, ch chan int, n int) {
    defer wg.Done()

    for i := 2; i <= n; i ++ {
        if i%2==0{
        ch <- i
        }
    }
}

func Odd(wg *sync.WaitGroup, ch chan int, n int) {
    defer wg.Done()

    for i := 1; i <= n; i ++ {
        if i%2!=0{
        ch <- i
        }
    }
}

func main() {
    n := 20
    ch := make(chan int)

    var wg sync.WaitGroup
    wg.Add(2)

    go Even(&wg, ch, n)
    go Odd(&wg, ch, n)

    go func() {
        wg.Wait()
        close(ch)
    }()

    for num := range ch {
        fmt.Println(num)
    }
}
